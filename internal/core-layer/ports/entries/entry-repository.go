package entries

import "gobanking/internal/core-layer/domain"

type EntryRepository interface {
	// given the account id, create an entry for this account
	Create(*domain.Entry) (*domain.Entry, error)

	// given the account id, get all entries of this account
	GetAll(int64) ([]*domain.Entry, error)

	// given the account id and the entry id, get the entry
	GetOne(int64, int64) (*domain.Entry, error)
}
