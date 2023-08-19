package models

import (
	"gobanking/internal/core-layer/domain"
	"time"
)

type PgTransfer struct {
	ID            int64     `json:"id"`
	ToAccountID   int64     `json:"to_account_id"`
	FromAccountID int64     `json:"from_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (pg_transfer *PgTransfer) ToDomainEntity() *domain.Transfer {
	return &domain.Transfer{
		ID:            pg_transfer.ID,
		ToAccountID:   pg_transfer.ToAccountID,
		FromAccountID: pg_transfer.FromAccountID,
		Amount:        pg_transfer.Amount,
		CreatedAt:     pg_transfer.CreatedAt,
		UpdatedAt:     pg_transfer.UpdatedAt,
	}
}
