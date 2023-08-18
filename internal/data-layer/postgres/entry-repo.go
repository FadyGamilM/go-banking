package postgres

import (
	"gobanking/internal/core-layer/domain"
	"gobanking/internal/data-layer/postgres/models"
	"gobanking/internal/infra-layer/db/postgres"
	"log"
)

type PG_EntryRepository struct {
	pg *postgres.PG_DB
}

const (
	create_Entry_For_Account_Query = `
		INSERT INTO entries 
		(account_id, amount)
		VALUES
		($1, $2)
		RETTURNING id, account_id, amount, created_at, updated_at
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
func (repo *PG_EntryRepository) CreateEntry(domain_entry *domain.Entry) (*domain.Entry, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	db_entry := new(models.PgEntry)

	err := repo.pg.DB.QueryRowContext(ctx, create_Entry_For_Account_Query, domain_entry.AccountID, domain_entry.Amount).Scan(&db_entry.ID, &db_entry.AccountID, &db_entry.Amount, &db_entry.CreatedAt, &db_entry.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	domain_entry = db_entry.ToDomainEntity()

	return domain_entry, nil
}

// given the account id, get all entries of this account
func (repo *PG_EntryRepository) GetAll(accID int64) ([]*domain.Entry, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	rows, err := repo.pg.DB.QueryContext(ctx, get_All_Entries_Of_Account_Query, accID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var domain_entries []*domain.Entry
	for rows.Next() {
		db_entry := new(models.PgEntry)
		err = rows.Scan(
			&db_entry.ID,
			&db_entry.AccountID,
			&db_entry.Amount,
			&db_entry.CreatedAt,
			&db_entry.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		domain_entries = append(domain_entries, db_entry.ToDomainEntity())
	}
	return domain_entries, nil
}

// given the account id and the entry id, get the entry
func (repo *PG_EntryRepository) GetOne(entryID int64, accID int64) (*domain.Entry, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	db_entry := new(models.PgEntry)
	err := repo.pg.DB.QueryRowContext(ctx, get_Entry_By_Entry_And_Account_IDs_Query, entryID, accID).Scan(&db_entry.ID, &db_entry.AccountID, &db_entry.Amount, &db_entry.CreatedAt, &db_entry.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db_entry.ToDomainEntity(), nil
}
