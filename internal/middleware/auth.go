package middleware

import (
	"errors"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"log"
)

func Auth(getUserHandler *queries.GetUserHandler, addUserHandler *commands.AddUserHandler) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			var u *user.User
			userIdDB, err := getUserHandler.Handle(uint(c.Sender().ID))

			log.Println(c.Sender().ID)
			log.Println(c.Message().Sender.ID)

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if userIdDB == nil {
				u = user.FromTelebotUser(c.Sender())

				//if err := addUserHandler.Handle(u); err != nil {
				//	return err
				//}
			} else {
				u = userIdDB
			}

			c.Set(string(user.UserInfoContextKey), u)

			return next(c)
		}
	}
}
