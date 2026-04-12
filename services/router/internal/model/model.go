package model

import "time"

// Account — аккаунт одного провайдера, принадлежащий пользователю
type Account struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	Provider     string    `db:"provider"`   // openai | claude | gemini
	EncryptedKey string    `db:"encrypted_key"`
	LimitPerDay  int       `db:"limit_per_day"`
	UsedToday    int       `db:"used_today"`
	ResetAt      time.Time `db:"reset_at"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
}

// RouteRequest — входящий запрос на маршрутизацию
type RouteRequest struct {
	UserID   string            `json:"user_id"`
	Provider string            `json:"provider"`
	Payload  map[string]any    `json:"payload"`
	Headers  map[string]string `json:"headers,omitempty"`
}

// RouteResult — результат маршрутизации
type RouteResult struct {
	AccountID  string `json:"account_id"`
	Provider   string `json:"provider"`
	StatusCode int    `json:"status_code"`
	Body       []byte `json:"body"`
}
