package account

import (
	"context"

	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"
)

type accountService struct {
	repo account.AccountRepo
}

func NewAccountService(r account.AccountRepo) account.AccountService {
	return &accountService{
		repo: r,
	}
}

func (s *accountService) Create(ctx context.Context, acc *account.Account) (*account.Account, error) {
	return s.repo.Create(ctx, acc)
}

func (s *accountService) GetAll(context.Context) ([]*account.Account, error) {
	return nil, nil
}

func (s *accountService) GetByID(context.Context, int64) (*account.Account, error) {
	return nil, nil
}

func (s *accountService) GetByOwnerName(ctx context.Context, ownerName string) (*account.Account, error) {
	return nil, nil
}

func (s *accountService) DeleteByID(ctx context.Context, accountID int64) error {
	return nil
}
func (s *accountService) UpdateByID(ctx context.Context, accountID int64, amount float64) (*account.Account, error) {
	return nil, nil
}
