package service

import "context"

// RouterService — интерфейс основной бизнес-логики роутера
type RouterService interface {
	// Route выбирает подходящий аккаунт и проксирует запрос
	Route(ctx context.Context, userID string, provider string, payload []byte) ([]byte, error)
	// GetAccountStatus возвращает текущий лимит аккаунтов пользователя
	GetAccountStatus(ctx context.Context, userID string) ([]AccountStatus, error)
}

type AccountStatus struct {
	AccountID   string `json:"account_id"`
	Provider    string `json:"provider"`
	Remaining   int    `json:"remaining"`
	ResetsAt    string `json:"resets_at"`
}
