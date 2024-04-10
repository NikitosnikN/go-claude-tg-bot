package handlers

import (
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/db"
	"gopkg.in/telebot.v3"
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
		text := c.Text()
		photo := c.Message().Photo
		if text == "" {
			text = imageDefaultPrompt // Ensure there's always text to accompany the photo
		}
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
			photo,
		)
	}
}
