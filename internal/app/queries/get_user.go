package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
)

type GetUserHandler struct {
	db *sql_store.SQLStore
}

func NewGetUserHandler(db *sql_store.SQLStore) *GetUserHandler {
	return &GetUserHandler{db: db}
}

func (h *GetUserHandler) Handle(userID uint) (*user.User, error) {
	u := &user.User{}

	result := h.db.DB().Where(&user.User{ID: userID}).First(u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}
