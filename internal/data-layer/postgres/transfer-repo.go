package postgres

import (
	"gobanking/internal/core-layer/domain"
	"gobanking/internal/data-layer/postgres/models"
	"gobanking/internal/infra-layer/db/postgres"
)

type PG_TransferRepository struct {
	pg *postgres.PG_DB
}

func NewPG_TransferRepo(pg *postgres.PG_DB) *PG_TransferRepository {
	return &PG_TransferRepository{
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

func (repo *PG_TransferRepository) Create(domain_transfer *domain.Transfer) (*domain.Transfer, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	db_transfer := new(models.PgTransfer)
	err := repo.pg.DB.QueryRowContext(ctx, create_Transfer_Query, domain_transfer.ToAccountID, domain_transfer.FromAccountID, domain_transfer.Amount).Scan(&db_transfer.ID, &db_transfer.ToAccountID, &db_transfer.FromAccountID, &db_transfer.Amount, &db_transfer.CreatedAt, &db_transfer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	domain_transfer = db_transfer.ToDomainEntity()
	return domain_transfer, nil
}
func (repo *PG_TransferRepository) GetAll() ([]*domain.Transfer, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	var domain_transfers []*domain.Transfer
	rows, err := repo.pg.DB.QueryContext(ctx, get_All_Transfers_Query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		db_transfer := new(models.PgTransfer)
		err = rows.Scan(
			&db_transfer.ID,
			&db_transfer.ToAccountID,
			&db_transfer.FromAccountID,
			&db_transfer.Amount,
			&db_transfer.CreatedAt,
			&db_transfer.UpdatedAt,
		)
		domain_transfers = append(domain_transfers, db_transfer.ToDomainEntity())
	}

	return domain_transfers, nil
}
func (repo *PG_TransferRepository) GetByID(id int64) (*domain.Transfer, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	db_transfer := new(models.PgTransfer)
	err := repo.pg.DB.QueryRowContext(ctx, get_Transfer_By_ID_Query, id).Scan(&db_transfer.ID, &db_transfer.ToAccountID, &db_transfer.FromAccountID, &db_transfer.Amount, &db_transfer.CreatedAt, &db_transfer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	domain_transfer := db_transfer.ToDomainEntity()
	return domain_transfer, nil
}
