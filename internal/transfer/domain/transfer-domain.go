package transfer

import (
	"context"
	"time"
)

type Transfer struct {
	ID            int64
	ToAccountID   int64
	FromAccountID int64
	Amount        float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransferRepository interface {
	Create(context.Context, *Transfer) (*Transfer, error)
	GetAll(context.Context) ([]*Transfer, error)
	GetByID(context.Context, int64) (*Transfer, error)
}
