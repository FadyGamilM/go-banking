package accounts

import "gobanking/internal/core-layer/domain"

type AccountRepository interface {
	Create(domain.Account) (*domain.Account, error)
}
