package ports

import "gobanking/internal/core/domain"

// => only the balance and the deleted fields are allowed to be modified, but we can't modify the account owner or the account currency type
type Update_account_params struct {
	Balance  *float64 // i made these types as * to the go-type because i need to check if these are provided or not
	Currency *string
}

type AccountRepository interface {
	Create(*domain.Account) error
	GetAll() ([]*domain.Account, error)
	GetById(int) (*domain.Account, error)
	GetByOwnerName(string) (*domain.Account, error)
	Update(int, Update_account_params) (*domain.Account, error)
	Delete(int) error
}
