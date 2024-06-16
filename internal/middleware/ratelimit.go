package middleware

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gopkg.in/telebot.v3"
)

// RateLimit middleware for tg bot
func RateLimit(requestsPerSecond int) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			u := c.Get(string(user.UserInfoContextKey)).(*user.User)

			if u == nil {
				return next(c)
			}

			// TODO implement rate limiter

			return next(c)
		}
	}
}
