package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	"gorm.io/gorm"
)

type AddMessagesToDialogHandler struct {
}

func NewAddMessagesToDialogHandler() *AddMessagesToDialogHandler {
	return &AddMessagesToDialogHandler{}
}

func (h *AddMessagesToDialogHandler) Handle(tx *gorm.DB, message *message.Message) error {
	result := tx.Create(message)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
