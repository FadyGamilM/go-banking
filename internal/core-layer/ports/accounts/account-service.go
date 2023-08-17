package accounts

import "gobanking/internal/common/types"

type AccountService interface {
	Create(*types.CreateAccountRequest) (types.CreateAccountResponse, error)
	GetAll() ([]*types.GetAccountResponse, error)
	GetById(int64) ([]*types.GetAccountResponse, error)
	DeleteById(int64) error
}
