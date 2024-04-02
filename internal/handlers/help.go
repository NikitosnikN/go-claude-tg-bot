package handlers

import "gopkg.in/telebot.v3"

func HelpHandler(c telebot.Context) error {
	text := `
\- Send a text message to get a completion
\- Send a photo to get an image summary
\- Send a photo and a text to get a completion
\- /echo \- echo message, debug purposes`

	return c.Send(text, telebot.ModeMarkdownV2)
}
