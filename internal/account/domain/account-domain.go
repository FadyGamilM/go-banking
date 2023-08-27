package account

import (
	"context"
	"time"
)

type Account struct {
	ID        int64
	OwnerName string
	Balance   float64
	Currency  string
	Activated bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountRepo interface {
	Create(context.Context, *Account) (*Account, error)
	GetAll(context.Context) ([]*Account, error)
	GetByID(context.Context, int64) (*Account, error)
	GetByOwnerName(context.Context, string) (*Account, error)
	DeleteByID(context.Context, int64) error
	UpdateByID(context.Context, int64, float64) (*Account, error)
}

type AccountService interface {
	Create(context.Context, *Account) (*Account, error)
	GetAll(context.Context) ([]*Account, error)
	GetByID(context.Context, int64) (*Account, error)
	GetByOwnerName(context.Context, string) (*Account, error)
	DeleteByID(context.Context, int64) error
	UpdateByID(context.Context, int64, float64) (*Account, error)
}
