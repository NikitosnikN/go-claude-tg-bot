package handlers

import (
	"context"
	anthropic "github.com/NikitosnikN/go-anthropic-api"
	"gopkg.in/telebot.v3"
	"io"
	"log"
)

const imageDefaultPrompt = "Analyze the image, describing its main subject, colors, composition, mood, and any symbolism present."

func PhotoMessageHandler(claude *anthropic.Client) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		var (
			text  = c.Text()
			bot   = c.Bot()
			photo = c.Message().Photo
			err   error
		)

		if text == "" {
			text = imageDefaultPrompt
		}

		if photo == nil {
			return c.Send("Sorry, cannot find a photo you have send to me. Please try again.")
		}

		reader, err := bot.File(photo.MediaFile())

		if err != nil {
			log.Printf("Cannot download photo: %v\n", err)
			err = c.Send("Sorry, cannot download the photo you have send to me. Please try again.")
			return err
		}

		payload, err := io.ReadAll(reader)

		if err != nil {
			log.Printf("Cannot read photo: %v\n", err)
			err = c.Send("Sorry, cannot read the photo you have send to me. Please try again.")
			return err
		}

		err = c.Notify(telebot.Typing)

		if err != nil {
			return err
		}

		// make request to Claude
		request := anthropic.NewMessageRequest("claude-3-haiku-20240307", 1024)
		request.AddImageMessage("user", payload, "image/jpeg", text)

		response, err := claude.CreateMessageRequest(context.Background(), *request)

		if err != nil {
			return err
		}

		return c.Send(response.Content[0].Text)
	}
}
