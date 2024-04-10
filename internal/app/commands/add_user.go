package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"gorm.io/gorm"
)

type AddUserHandler struct {
}

func NewAddUserHandler() *AddUserHandler {
	return &AddUserHandler{}
}

func (h *AddUserHandler) Handle(tx *gorm.DB, user *user.User) error {

	result := tx.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
