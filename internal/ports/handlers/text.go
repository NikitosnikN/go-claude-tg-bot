package handlers

import (
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/db"
	"gopkg.in/telebot.v3"
)

func TextMessageHandler(
	claude *anthropic.Client,
	newTx db.NewDBTx,
	getLatestDialog *queries.GetLatestDialogHandler,
	getDialogMessages *queries.GetDialogMessagesHandler,
	getPhoto *queries.GetPhoto,
	addDialog *commands.AddDialogHandler,
	addMessageToDialog *commands.AddMessagesToDialogHandler,

) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		text := c.Text()
		return HandleDialogInteraction(
			c,
			claude,
			newTx,
			getLatestDialog,
			getDialogMessages,
			getPhoto,
			addDialog,
			addMessageToDialog,
			text,
			nil, // No photo in text handler
		)
	}
}
