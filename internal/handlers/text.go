package handlers

import (
	"context"
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"gopkg.in/telebot.v3"
)

func TextMessageHandler(claude *anthropic.Client) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		var (
			text = c.Text()
		)
		err := c.Notify(telebot.Typing)

		if err != nil {
			return err
		}

		request := anthropic.NewMessageRequest("claude-3-haiku-20240307", 1024)
		request.AddTextMessage("user", text)

		response, err := claude.CreateMessageRequest(context.Background(), *request)

		if err != nil {
			return err
		}

		return c.Send(response.Content[0].Text)
	}
}
