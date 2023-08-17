package usecases

import (
	"gobanking/internal/common/types"
	"gobanking/internal/data-layer/store"
)

type accountService struct {
	Data *store.DataStore
}

type Config struct {
	Data *store.DataStore
}

// Create a new account service instance
func NewAccountService(c *Config) *accountService {
	return &accountService{
		Data: c.Data,
	}
}

func (as *accountService) Create(req *types.CreateAccountRequest) (types.CreateAccountResponse, error) {
}

func (as *accountService) GetAll() ([]*types.GetAccountResponse, error)          { return nil, nil }
func (as *accountService) GetById(id int64) ([]*types.GetAccountResponse, error) {}
func (as *accountService) DeleteById(id int64) error                             {}
