package entry

import (
	"context"
	"log"

	"github.com/FadyGamilM/go-banking-v2/db"
	domain "github.com/FadyGamilM/go-banking-v2/internal/entry/domain"
)

type entryRepo struct {
	pg *db.PG
}

func NewEntryRepo(pg *db.PG) *entryRepo {
	return &entryRepo{
		pg: pg,
	}
}

const (
	create_Entry_For_Account_Query = `
		INSERT INTO entries 
		(account_id, amount)
		VALUES
		($1, $2)
		RETURNING id, account_id, amount, created_at, updated_at
	`

	get_All_Entries_Of_Account_Query = `
		SELECT id, account_id, amount, created_at, updated_at 
		FROM entries 
		WHERE account_id = $1	
	`

	get_Entry_By_Entry_And_Account_IDs_Query = `
		SELECT id, account_id, amount, created_at, updated_at 
		FROM entries 
		WHERE id = $1 AND account_id = $2
	`
)

// given the account id, create an entry for this account
func (repo *entryRepo) Create(ctx context.Context, entry *domain.Entry) (*domain.Entry, error) {

	domainEntry := new(domain.Entry)

	err := repo.pg.DB.QueryRowContext(ctx, create_Entry_For_Account_Query, entry.AccountID, entry.Amount).Scan(
		&domainEntry.ID,
		&domainEntry.AccountID,
		&domainEntry.Amount,
		&domainEntry.CreatedAt,
		&domainEntry.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domainEntry, nil
}

// given the account id, get all entries of this account
func (repo *entryRepo) GetAll(ctx context.Context, accID int64) ([]*domain.Entry, error) {
	rows, err := repo.pg.DB.QueryContext(ctx, get_All_Entries_Of_Account_Query, accID)
	if err != nil {
		return nil, err
	}

	var domainEntries []*domain.Entry
	for rows.Next() {
		entry := new(domain.Entry)
		err = rows.Scan(
			&entry.ID,
			&entry.AccountID,
			&entry.Amount,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		domainEntries = append(domainEntries, entry)
	}
	return domainEntries, nil
}

// given the account id and the entry id, get the entry
func (repo *entryRepo) GetOne(ctx context.Context, entryID int64, accID int64) (*domain.Entry, error) {

	entry := new(domain.Entry)
	err := repo.pg.DB.QueryRowContext(ctx, get_Entry_By_Entry_And_Account_IDs_Query, entryID, accID).Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return entry, nil
}
