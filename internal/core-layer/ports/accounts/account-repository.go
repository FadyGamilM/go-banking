package accounts

import "gobanking/internal/core-layer/domain"

type AccountRepository interface {
	Create(*domain.Account) (*domain.Account, error)
	GetAll() ([]*domain.Account, error)
	GetByID(int64) (*domain.Account, error)
	GetByOwnerName(string) (*domain.Account, error)
	DeleteByID(int64) error
}
