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
	"github.com/NikitosnikN/go-claude-tg-bot/internal/utils"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
)

const imageDefaultPrompt = "Analyze the image, describing its main subject, colors, composition, mood, and any symbolism present."

func PhotoMessageHandler(
	claude *anthropic.Client,
	newTx db.NewDBTx,
	getLatestDialog *queries.GetLatestDialogHandler,
	getDialogMessages *queries.GetDialogMessagesHandler,
	getPhoto *queries.GetPhoto,
	addDialog *commands.AddDialogHandler,
	addMessageToDialog *commands.AddMessagesToDialogHandler,

) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("TextMessageHandler panic: %v", r)
			}
		}()

		tx := newTx()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		text := c.Text()
		photo := c.Message().Photo
		u := c.Get(string(user.UserInfoContextKey)).(*user.User)
		var err error
		var d *dialog.Dialog
		var messages []*message.Message

		if photo == nil {
			return c.Send("Sorry, cannot find a photo you have send to me. Please try again.")
		}

		if text == "" {
			text = imageDefaultPrompt
		}

		d, err = getLatestDialog.Handle(tx, u.ID)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			d = dialog.NewDialog(u.ID)

			err = addDialog.Handle(tx, d)

			if err != nil {
				tx.Rollback()
				return err
			}
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

		// add new photo message
		photoPayload, err := getPhoto.Handle(photo.FileID)

		if err != nil {
			tx.Rollback()
			return err
		}

		request.AddImageMessage("user", photoPayload, "image/jpeg", text)

		response, err := claude.CreateMessageRequest(context.Background(), *request)

		if err != nil {
			return err
		}

		answer := response.Content[0].Text

		// save user message to DB
		err = addMessageToDialog.Handle(tx, message.NewMessage(d.ID, "user", text, photo.FileID))
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
}
