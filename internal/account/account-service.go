package account

import (
	"context"
	"database/sql"
	"log"

	"github.com/FadyGamilM/go-banking-v2/common"
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

func (s *accountService) GetByID(ctx context.Context, id int64) (*account.Account, error) {
	log.Println("the id at the service layer is : ", id)
	acc, err := s.repo.GetByID(ctx, id)
	if err != nil {
		// this should be handled by a proper error handler "constant string"
		if err == sql.ErrNoRows {
			return nil, common.NotFound
		}
		return nil, common.InternalServerError
	}

	return acc, nil
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

func (s *accountService) GetAllInPages(ctx context.Context, limit, offset int32) ([]*account.Account, error) {
	return s.repo.GetAllInPages(ctx, limit, offset)
}
