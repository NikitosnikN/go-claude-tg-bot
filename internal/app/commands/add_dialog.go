package commands

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
)

type AddDialogHandler struct {
	db *sql_store.SQLStore
}

func NewAddDialogHandler(db *sql_store.SQLStore) *AddDialogHandler {
	return &AddDialogHandler{
		db: db,
	}
}

func (h *AddDialogHandler) Handle(dialog *dialog.Dialog) error {
	result := h.db.DB().Create(dialog)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
