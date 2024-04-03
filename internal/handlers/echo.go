package handlers

import "gopkg.in/telebot.v3"

// EchoHandler debug handler
func EchoHandler(c telebot.Context) error {
	var (
		text = c.Text()
	)
	return c.Send(text, telebot.ModeMarkdownV2)
}
