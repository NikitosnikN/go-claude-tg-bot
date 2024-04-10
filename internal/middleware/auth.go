package middleware

import (
	"errors"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/queries"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/db"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

func Auth(newTx db.NewDBTx, getUserHandler *queries.GetUserHandler, addUserHandler *commands.AddUserHandler) telebot.MiddlewareFunc {
	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			var u *user.User
			tx := newTx()

			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			userIdDB, err := getUserHandler.Handle(tx, uint(c.Sender().ID))

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if userIdDB == nil {
				u = user.FromTelebotUser(c.Sender())

				if err = addUserHandler.Handle(tx, u); err != nil {
					tx.Rollback()
					return err
				}
			} else {
				u = userIdDB
			}

			c.Set(string(user.UserInfoContextKey), u)

			tx.Commit()

			return next(c)
		}
	}
}
