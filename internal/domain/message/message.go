package message

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/dialog"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        uuid.UUID
	DialogID  uuid.UUID
	Dialog    *dialog.Dialog
	Role      string
	Text      string
	PhotoID   string
	CreatedAt time.Time `gorm:"index"`
}

func NewMessage(DialogID uuid.UUID, Role string, Text string, PhotoID string) *Message {
	return &Message{
		ID:        uuid.New(),
		DialogID:  DialogID,
		Role:      Role,
		Text:      Text,
		PhotoID:   PhotoID,
		CreatedAt: time.Now(),
	}
}
