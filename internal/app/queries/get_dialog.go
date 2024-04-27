package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	aerrors "github.com/NikitosnikN/go-claude-tg-bot/internal/errors"
	"gorm.io/gorm"
	"time"
)

type GetLatestDialogHandler struct{}

func NewGetLatestDialogHandler() *GetLatestDialogHandler {
	return &GetLatestDialogHandler{}
}

func (h *GetLatestDialogHandler) Handle(tx *gorm.DB, userID uint, interval time.Duration) (*dialog.Dialog, error) {
	d := &dialog.Dialog{}

	result := tx.Order("created_at desc").Where(&dialog.Dialog{UserID: userID}).First(d)

	if result.Error != nil {
		return nil, result.Error
	}

	m := &message.Message{}

	result = tx.Where(&message.Message{DialogID: d.ID}).Order("created_at desc").First(m)

	if result.Error != nil {
		return nil, result.Error
	}

	if time.Since(m.CreatedAt) > interval {
		return nil, aerrors.DialogExpired
	}

	return d, nil
}
