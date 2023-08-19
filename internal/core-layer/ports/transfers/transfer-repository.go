package transfers

import "gobanking/internal/core-layer/domain"

type TransferRepository interface {
	Create(*domain.Transfer) (*domain.Transfer, error)
	GetAll() ([]*domain.Transfer, error)
	GetByID(int64) (*domain.Transfer, error)
}
