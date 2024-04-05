package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
)

type AddUserHandler struct {
	db *sql_store.SQLStore
}

func NewAddUserHandler(db *sql_store.SQLStore) *AddUserHandler {
	return &AddUserHandler{
		db: db,
	}
}

func (h *AddUserHandler) Handle(user *user.User) error {

	result := h.db.DB().Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
