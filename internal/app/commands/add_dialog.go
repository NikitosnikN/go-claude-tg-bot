package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"gorm.io/gorm"
)

type AddDialogHandler struct {
}

func NewAddDialogHandler() *AddDialogHandler {
	return &AddDialogHandler{}
}

func (h *AddDialogHandler) Handle(tx *gorm.DB, dialog *dialog.Dialog) error {
	result := tx.Create(dialog)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
