package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
	"github.com/google/uuid"
)

type GetDialogMessagesHandler struct {
	db *sql_store.SQLStore
}

func NewGetDialogMessagesHandler(db *sql_store.SQLStore) *GetDialogMessagesHandler {
	return &GetDialogMessagesHandler{db: db}
}

func (h *GetDialogMessagesHandler) Handle(dialogID uuid.UUID) ([]*message.Message, error) {
	var messages []*message.Message

	result := h.db.DB().Order("created_at asc").Where(&message.Message{DialogID: dialogID}).Find(&messages)

	if result.Error != nil {
		return nil, result.Error
	}

	return messages, nil
}
