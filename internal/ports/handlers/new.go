package handlers

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/app/commands"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/db"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gopkg.in/telebot.v3"
	"log"
)

// NewHandler start new dialog
func NewHandler(
	newTx db.NewDBTx,
	addDialog *commands.AddDialogHandler,

) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("NewHandler panic: %v", r)
			}
		}()

		tx := newTx()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		u := c.Get(string(user.UserInfoContextKey)).(*user.User)
		err := addDialog.Handle(tx, dialog.NewDialog(u.ID))

		if err != nil {
			tx.Rollback()
			return err
		}

		err = c.Send("Hi, how can I help you?")

		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()

		return err
	}
}
