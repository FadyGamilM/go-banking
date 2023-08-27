package transfer

import (
	"context"

	transfer "github.com/FadyGamilM/go-banking-v2/internal/transfer/domain"
)

type transferService struct {
	repo transfer.TransferService
}

func NewTransferService(r transfer.TransferRepo) transfer.TransferService {
	return &transferService{
		repo: r,
	}
}

// given the account id, create an entry for this account
func (s *transferService) Create(context.Context, *transfer.Transfer) (*transfer.Transfer, error) {
	return nil, nil
}

// given the account id, get all entries of this account
func (s *transferService) GetAll(context.Context) ([]*transfer.Transfer, error) {
	return nil, nil
}

// given the account id and the entry id, get the entry
func (s *transferService) GetByID(context.Context, int64) (*transfer.Transfer, error) {
	return nil, nil
}
