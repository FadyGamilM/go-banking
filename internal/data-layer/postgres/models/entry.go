package models

import (
	"gobanking/internal/core-layer/domain"
	"time"
)

type PgEntry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (pg_entry *PgEntry) ToDomainEntity() *domain.Entry {
	return &domain.Entry{
		ID:        pg_entry.ID,
		AccountID: pg_entry.AccountID,
		Amount:    pg_entry.Amount,
		CreatedAt: pg_entry.CreatedAt,
		UpdatedAt: pg_entry.UpdatedAt,
	}
}
