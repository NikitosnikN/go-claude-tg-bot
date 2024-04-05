package queries

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/adapter/sql_store"
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
)

type GetLatestDialogHandler struct {
	db *sql_store.SQLStore
}

func NewGetLatestDialogHandler(db *sql_store.SQLStore) *GetLatestDialogHandler {
	return &GetLatestDialogHandler{db: db}
}

func (h *GetLatestDialogHandler) Handle(userID uint) (*dialog.Dialog, error) {
	u := &dialog.Dialog{}

	result := h.db.DB().Order("created_at desc").Where(&dialog.Dialog{UserID: userID}).First(u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}
