package handlers

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/utils"
	"gopkg.in/telebot.v3"
)

func HelpHandler(c telebot.Context) error {
	text := `
- Send a text message to get a completion.
- Send a photo to get an image summary.
- Send a photo and a text to get a completion.
- /new - start new dialog.`
	return c.Send(utils.EscapeMarkdown(text), telebot.ModeMarkdownV2)
}
