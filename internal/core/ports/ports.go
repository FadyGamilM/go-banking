package ports

import "gobanking/internal/core/domain"

// => only the balance and the deleted fields are allowed to be modified, but we can't modify the account owner or the account currency type
type Update_account_params struct {
	Balance  *float64 // i made these types as * to the go-type because i need to check if these are provided or not
	Currency *string
}

// Now the core.service layer don't need to know anything about the impl, it will just depends on this interface
type AccountRepository interface {
	Create(*domain.Account) error
	GetAll() ([]*domain.Account, error)
	GetById(int) (*domain.Account, error)
	GetByOwnerName(string) (*domain.Account, error)
	Update(int, Update_account_params) (*domain.Account, error)
	Delete(int) error
}

type EntryRepository interface {
	Create(*domain.Entry) error
	GetAll() ([]*domain.Entry, error)
	GetById(int) (*domain.Entry, error)
	GetByAccountId(int) (*domain.Entry, error)
	Delete(int) error
	DeleteByAccountId(int) error
}

type TransferRepository interface {
	Create(*domain.Transfer) error
	GetAll() ([]*domain.Transfer, error)
	GetById(int) (*domain.Transfer, error)
	GetByToAccountId(int) (*domain.Transfer, error)
	GetByFromAccountId(int) (*domain.Transfer, error)
	GetByAccountId(int) (*domain.Transfer, error)
	Delete(int) error
}