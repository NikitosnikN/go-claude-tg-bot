package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"gorm.io/gorm"
)

type GetLatestDialogHandler struct{}

func NewGetLatestDialogHandler() *GetLatestDialogHandler {
	return &GetLatestDialogHandler{}
}

func (h *GetLatestDialogHandler) Handle(tx *gorm.DB, userID uint) (*dialog.Dialog, error) {
	u := &dialog.Dialog{}

	result := tx.Order("created_at desc").Where(&dialog.Dialog{UserID: userID}).First(u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}
