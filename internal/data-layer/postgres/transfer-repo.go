package postgres

import "gobanking/internal/infra-layer/db/postgres"

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
	`
)
