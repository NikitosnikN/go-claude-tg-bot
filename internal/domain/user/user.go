package user

import (
	"github.com/NikitosnikN/go-claude-tg-bot/internal/domain/claude_model"
	"gopkg.in/telebot.v3"
)

type contextKey string

const (
	UserInfoContextKey contextKey = "userInfo"
)

type User struct {
	ID          uint
	Username    string
	FirstName   string
	LastName    string
	ClaudeModel string
}

func FromTelebotUser(user *telebot.User) *User {
	return &User{
		ID:          uint(user.ID),
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		ClaudeModel: string(claude_model.ClaudeDefaultModel),
	}
}
