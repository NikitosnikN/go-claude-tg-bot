package middleware

import (
	"gopkg.in/telebot.v3"
)

func ErrorHandler() telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			originError := next(c)
			if originError != nil {
				_ = c.Send("Got error while proceeding message, try again or check logs")
			}
			return originError
		}
	}
}
