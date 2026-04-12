package repository

import "context"

// SessionRepository — работа с сессиями
type SessionRepository interface {
	Create(ctx context.Context, userID, provider string) (string, error)
	FindByID(ctx context.Context, id string) (*Session, error)
	Delete(ctx context.Context, id string) error
}

// MessageRepository — работа с историей сообщений
type MessageRepository interface {
	Save(ctx context.Context, sessionID, role, content string, tokenCount int) error
	// GetBySession возвращает сообщения от новейших к старейшим с лимитом
	GetBySession(ctx context.Context, sessionID string, limit int) ([]Message, error)
}

type Session struct {
	ID       string
	UserID   string
	Provider string
}

type Message struct {
	ID         string
	SessionID  string
	Role       string
	Content    string
	TokenCount int
}
