package service

import "context"

// ContextService — интерфейс управления контекстом сессий
type ContextService interface {
	// SaveMessage сохраняет сообщение в историю сессии
	SaveMessage(ctx context.Context, sessionID, role, content string) error
	// GetContext возвращает историю сессии, готовую к передаче в API
	// maxTokens — лимит токенов нового аккаунта (обрезает историю если нужно)
	GetContext(ctx context.Context, sessionID string, maxTokens int) (*ContextSnapshot, error)
	// CreateSession создаёт новую сессию
	CreateSession(ctx context.Context, userID, provider string) (string, error)
	// DeleteSession удаляет сессию и всю историю
	DeleteSession(ctx context.Context, sessionID string) error
}

type ContextSnapshot struct {
	SessionID   string    `json:"session_id"`
	Messages    []Message `json:"messages"`
	TotalTokens int       `json:"total_tokens"`
	Truncated   bool      `json:"truncated"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
