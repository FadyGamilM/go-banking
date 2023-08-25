package transfer

import (
	"context"

	"github.com/FadyGamilM/go-banking-v2/db"
	domain "github.com/FadyGamilM/go-banking-v2/internal/transfer/domain"
)

type transferRepo struct {
	pg *db.PG
}

func NewTransferRepo(pg *db.PG) *transferRepo {
	return &transferRepo{
		pg: pg,
	}
}

const (
	create_Transfer_Query = `
		INSERT INTO transfers
		(to_account, from_account, amount)
		VALUES
		($1, $2, $3)
		RETURNING id, to_account, from_account, amount, created_at, updated_at
	`

	get_All_Transfers_Query = `
		SELECT id, to_account, from_account, amount, created_at, updated_at
		FROM transfers
	`

	get_Transfer_By_ID_Query = `
		SELECT id, to_account, from_account, amount, created_at, updated_at
		FROM transfers 
		WHERE id = $1
	`
)

func (repo *transferRepo) Create(ctx context.Context, transfer *domain.Transfer) (*domain.Transfer, error) {
	domainTransfer := new(domain.Transfer)
	err := repo.pg.DB.QueryRowContext(ctx, create_Transfer_Query, transfer.ToAccountID, transfer.FromAccountID, transfer.Amount).Scan(
		&domainTransfer.ID,
		&domainTransfer.ToAccountID,
		&domainTransfer.FromAccountID,
		&domainTransfer.Amount,
		&domainTransfer.CreatedAt,
		&domainTransfer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domainTransfer, nil
}

func (repo *transferRepo) GetAll(ctx context.Context) ([]*domain.Transfer, error) {

	var transfers []*domain.Transfer
	rows, err := repo.pg.DB.QueryContext(ctx, get_All_Transfers_Query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transfer := new(domain.Transfer)
		err = rows.Scan(
			&transfer.ID,
			&transfer.ToAccountID,
			&transfer.FromAccountID,
			&transfer.Amount,
			&transfer.CreatedAt,
			&transfer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (repo *transferRepo) GetByID(ctx context.Context, id int64) (*domain.Transfer, error) {

	transfer := new(domain.Transfer)
	err := repo.pg.DB.QueryRowContext(ctx, get_Transfer_By_ID_Query, id).Scan(
		&transfer.ID,
		&transfer.ToAccountID,
		&transfer.FromAccountID,
		&transfer.Amount,
		&transfer.CreatedAt,
		&transfer.UpdatedAt,
	)
	if err != nil {
		return &domain.Transfer{}, err
	}

	return transfer, nil
}
