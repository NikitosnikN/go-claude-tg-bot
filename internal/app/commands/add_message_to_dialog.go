package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/message"
)

type AddMessagesToDialogHandler struct {
	db *sql_store.SQLStore
}

func NewAddMessagesToDialogHandler(db *sql_store.SQLStore) *AddMessagesToDialogHandler {
	return &AddMessagesToDialogHandler{
		db: db,
	}
}

func (h *AddMessagesToDialogHandler) Handle(message *message.Message) error {
	result := h.db.DB().Create(message)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
