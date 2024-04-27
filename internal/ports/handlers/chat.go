package handlers

import (
	"context"
	"errors"
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/db"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	aerrors "github.com/NikitosnikN/go-claude-tg-bot/internal/errors"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/utils"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
	"time"
)

const lastMessageExpireInteval = 15 * time.Minute

func HandleDialogInteraction(
	c telebot.Context,
	claude *anthropic.Client,
	newTx db.NewDBTx,
	getLatestDialog *queries.GetLatestDialogHandler,
	getDialogMessages *queries.GetDialogMessagesHandler,
	getPhoto *queries.GetPhoto,
	addDialog *commands.AddDialogHandler,
	addMessageToDialog *commands.AddMessagesToDialogHandler,
	text string,
	photo *telebot.Photo,
) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("HandleDialogInteraction panic: %v", r)
		}
	}()

	tx := newTx()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	u := c.Get(string(user.UserInfoContextKey)).(*user.User)
	var err error
	var d *dialog.Dialog
	var messages []*message.Message

	d, err = getLatestDialog.Handle(tx, u.ID, lastMessageExpireInteval)

	if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, aerrors.DialogExpired) {
		// create new dialog
		d = dialog.NewDialog(u.ID)

		err = addDialog.Handle(tx, d)

		if err != nil {
			tx.Rollback()
			return err
		}

		err = c.Send("Starting new dialog...")

		if err != nil {
			tx.Rollback()
			return err
		}
	} else if err != nil {
		return err
	} else {
		messages, err = getDialogMessages.Handle(tx, d.ID)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	err = c.Notify(telebot.Typing)

	if err != nil {
		return err
	}

	request := anthropic.NewMessageRequest(u.ClaudeModel, 1024)

	// add old messages from history
	for _, m := range messages {
		if m.PhotoID != "" {
			photoPayload, err := getPhoto.Handle(m.PhotoID)
			if err != nil {
				tx.Rollback()
				return err
			}
			request.AddImageMessage(anthropic.MessageRole(m.Role), photoPayload, "image/jpeg", m.Text)
		} else {
			request.AddTextMessage(anthropic.MessageRole(m.Role), m.Text)
		}
	}

	// add new message or photo
	if photo != nil {
		photoPayload, err := getPhoto.Handle(photo.FileID)
		if err != nil {
			tx.Rollback()
			return err
		}
		request.AddImageMessage("user", photoPayload, "image/jpeg", text)
	} else {
		request.AddTextMessage("user", text)
	}

	response, err := claude.CreateMessageRequest(context.Background(), *request)

	if err != nil {
		return err
	}

	answer := response.Content[0].Text

	// save user message to DB
	if photo != nil {
		err = addMessageToDialog.Handle(tx, message.NewMessage(d.ID, "user", text, photo.FileID))
	} else {
		err = addMessageToDialog.Handle(tx, message.NewMessage(d.ID, "user", text, ""))
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	// save assistant message to DB
	err = addMessageToDialog.Handle(tx, message.NewMessage(d.ID, "assistant", answer, ""))
	if err != nil {
		tx.Rollback()
		return err
	}

	err = c.Send(utils.EscapeMarkdown(answer), telebot.ModeMarkdownV2)

	if err != nil {
		return err
	}

	tx.Commit()
	return err
}
