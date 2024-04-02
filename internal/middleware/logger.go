package middleware

import (
	"gopkg.in/telebot.v3"
	"log"
)

func VerboseLogger() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			username := c.Sender().Username
			log.Printf("New message from %s", username)
			return next(c)
		}
	}
}
