package entry

import (
	"context"
	"time"
)

type Entry struct {
	ID        int64
	AccountID int64
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EntryRepo interface {
	// given the account id, create an entry for this account
	Create(context.Context, *Entry) (*Entry, error)

	// given the account id, get all entries of this account
	GetAll(context.Context, int64) ([]*Entry, error)

	// given the account id and the entry id, get the entry
	GetOne(context.Context, int64, int64) (*Entry, error)
}

type EntryService interface {
	// given the account id, create an entry for this account
	Create(context.Context, *Entry) (*Entry, error)

	// given the account id, get all entries of this account
	GetAll(context.Context, int64) ([]*Entry, error)

	// given the account id and the entry id, get the entry
	GetOne(context.Context, int64, int64) (*Entry, error)
}
