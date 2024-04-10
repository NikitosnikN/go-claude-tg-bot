package queries

import (
	"gopkg.in/telebot.v3"
	"io"
)

type GetPhoto struct {
	bot *telebot.Bot
}

func NewGetPhoto(bot *telebot.Bot) *GetPhoto {
	return &GetPhoto{
		bot: bot,
	}
}

func (h *GetPhoto) Handle(photoId string) ([]byte, error) {

	file, err := h.bot.FileByID(photoId)

	if err != nil {
		return []byte{}, err
	}

	reader, err := h.bot.File(&file)

	if err != nil {
		return []byte{}, err
	}

	payload, err := io.ReadAll(reader)

	if err != nil {
		return []byte{}, err
	}

	return payload, nil
}
