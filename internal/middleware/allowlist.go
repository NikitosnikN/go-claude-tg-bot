package middleware

import (
	"gopkg.in/telebot.v3"
	"log"
	"slices"
	"strings"
)

func AllowList(usernames ...string) telebot.MiddlewareFunc {
	usernamesLowerCase := make([]string, len(usernames))
	for _, v := range usernames {
		usernamesLowerCase = append(usernamesLowerCase, strings.ToLower(v))
	}

	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			username := strings.ToLower(c.Sender().Username)

			if len(usernamesLowerCase) != 0 && slices.Index(usernamesLowerCase, username) == -1 {
				log.Printf("User %s is not allowed to use this bot.\n", username)
				return c.Send("You are not allowed to use this bot.")
			}

			return next(c)
		}
	}
}
