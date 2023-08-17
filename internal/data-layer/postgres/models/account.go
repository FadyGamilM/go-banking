package models

import (
	"gobanking/internal/core-layer/domain"
	"time"
)

type PgAccount struct {
	ID        int64     `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Activated bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (pg_acc *PgAccount) ToDomainEntity() *domain.Account {
	acc := new(domain.Account)
	acc.ID = pg_acc.ID
	acc.OwnerName = pg_acc.OwnerName
	acc.Balance = pg_acc.Balance
	acc.Currency = pg_acc.Currency
	acc.Activated = pg_acc.Activated
	acc.CreatedAt = pg_acc.CreatedAt
	acc.UpdatedAt = pg_acc.UpdatedAt

	return acc
}
