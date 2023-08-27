package entry

import (
	"context"

	entry "github.com/FadyGamilM/go-banking-v2/internal/entry/domain"
)

type entryService struct {
	repo entry.EntryRepo
}

func NewEntryService(r entry.EntryRepo) entry.EntryService {
	return &entryService{
		repo: r,
	}
}

// given the account id, create an entry for this account
func (s *entryService) Create(context.Context, *entry.Entry) (*entry.Entry, error) {
	return nil, nil
}

// given the account id, get all entries of this account
func (s *entryService) GetAll(context.Context, int64) ([]*entry.Entry, error) {
	return nil, nil
}

// given the account id and the entry id, get the entry
func (s *entryService) GetOne(context.Context, int64, int64) (*entry.Entry, error) {
	return nil, nil
}
