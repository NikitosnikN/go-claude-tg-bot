package handlers

import "gopkg.in/telebot.v3"

func StartHandler(c telebot.Context) error {
	text := `Hi, how can I help you?

If you have any questions related to bot usage, please use command /help`

	return c.Send(text)
}
