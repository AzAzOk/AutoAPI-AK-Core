package model

import "time"

// Session — сессия пользователя с конкретным провайдером
type Session struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id"`
	Provider  string    `db:"provider" json:"provider"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// Message — одно сообщение в сессии
type Message struct {
	ID        string    `db:"id" json:"id"`
	SessionID string    `db:"session_id" json:"session_id"`
	Role      string    `db:"role" json:"role"`   // user | assistant | system
	Content   string    `db:"content" json:"content"`
	TokenCount int      `db:"token_count" json:"token_count"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// ContextSnapshot — готовый контекст для передачи в API провайдера
// Формируется при смене аккаунта
type ContextSnapshot struct {
	SessionID string    `json:"session_id"`
	Messages  []Message `json:"messages"`
	TotalTokens int     `json:"total_tokens"`
	Truncated bool      `json:"truncated"` // true если история была обрезана
}
