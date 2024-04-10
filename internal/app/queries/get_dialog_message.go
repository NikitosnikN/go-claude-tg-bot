package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetDialogMessagesHandler struct{}

func NewGetDialogMessagesHandler() *GetDialogMessagesHandler {
	return &GetDialogMessagesHandler{}
}

func (h *GetDialogMessagesHandler) Handle(tx *gorm.DB, dialogID uuid.UUID) ([]*message.Message, error) {
	var messages []*message.Message

	result := tx.Order("created_at asc").Where(&message.Message{DialogID: dialogID}).Find(&messages)

	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
