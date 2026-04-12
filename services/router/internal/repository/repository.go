package repository

import "context"

// AccountRepository — интерфейс работы с хранилищем аккаунтов
type AccountRepository interface {
	// FindAvailable ищет аккаунт с доступным лимитом для провайдера
	FindAvailable(ctx context.Context, userID, provider string) (*Account, error)
	// IncrementUsage увеличивает счётчик использования на 1
	IncrementUsage(ctx context.Context, accountID string) error
	// GetAllByUser возвращает все аккаунты пользователя
	GetAllByUser(ctx context.Context, userID string) ([]Account, error)
}

type Account struct {
	ID          string
	Provider    string
	APIKey      string // расшифрованный ключ (только в памяти)
	Remaining   int
}
