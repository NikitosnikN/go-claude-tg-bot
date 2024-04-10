package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gorm.io/gorm"
)

type GetUserHandler struct {
}

func NewGetUserHandler() *GetUserHandler {
	return &GetUserHandler{}
}

func (h *GetUserHandler) Handle(tx *gorm.DB, userID uint) (*user.User, error) {
	u := &user.User{}

	result := tx.Where(&user.User{ID: userID}).First(u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}
