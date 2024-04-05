package dialog

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/user"
	"github.com/google/uuid"
	"time"
)

type Dialog struct {
	ID        uuid.UUID
	UserID    uint
	User      *user.User
	CreatedAt time.Time `gorm:"index"`
}

func NewDialog(UserID uint) *Dialog {
	return &Dialog{
		ID:        uuid.New(),
		UserID:    UserID,
		CreatedAt: time.Now(),
	}
}
