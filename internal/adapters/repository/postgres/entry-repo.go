package postgres

import (
	"gobanking/internal/core/domain"
	"log"
)

type Entry_repo struct {
	repo *Postgres_repo
}

func (er *Entry_repo) Create(entry *domain.Entry) error {
	// define context
	ctx, cancel := CreateContext()
	defer cancel()

	query := `
		INSERT INTO entries (account_id, amount)
		VALUES $1, $2
	`

	_, err := er.repo.db.ExecContext(ctx, query, entry.AccountID, entry.Amount)
	if err != nil {
		log.Println("error during the execution of the INSERT query at the repository level : %v \n", err)
		return err
	}

	return nil
}

func (er *Entry_repo) GetAll() ([]*domain.Entry, error) {
	// define context
	ctx, cancel := CreateContext()
	defer cancel()

	query := `SELECT * FROM entries`

	rows, err := er.repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("error during the execution of the SELECT query at the repository level : %v \n", err)
		return nil, err
	}

	var entries []*domain.Entry
	for rows.Next() {
		var entry *domain.Entry
		err := rows.Scan(
			&entry.ID,
			&entry.AccountID,
			&entry.Amount,
			&entry.CreatedAt,
			&entry.UpdatedAt,
		)

		if err != nil {
			log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (er *Entry_repo) GetById(id int) (*domain.Entry, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `SELECT * FROM entries WHERE id = $1`

	row := er.repo.db.QueryRowContext(ctx, query, id)

	var entry *domain.Entry
	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
		return nil, err
	}
	return entry, nil
}

func (er *Entry_repo) GetByAccountId(account_id int) (*domain.Entry, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `SELECT * FROM entries WHERE account_id = $1`

	row := er.repo.db.QueryRowContext(ctx, query, account_id)

	var entry *domain.Entry
	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
		&entry.UpdatedAt,
	)
	if err != nil {
		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
		return nil, err
	}
	return entry, nil
}

func (er *Entry_repo) Delete(id int) error {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `DELETE FROM entries WHERE id = $1`

	_, err := er.repo.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("error during the execution of the DELETE query inside postgres repository layer : %v \n", err)
		return err
	}
	return nil
}

func (er *Entry_repo) DeleteByAccountId(account_id int) error {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `DELETE FROM entries WHERE account_id = $1`

	_, err := er.repo.db.ExecContext(ctx, query, account_id)
	if err != nil {
		log.Printf("error during the execution of the DELETE query inside postgres repository layer : %v \n", err)
		return err
	}
	return nil
}

func NewEntryRepo(repo *Postgres_repo) *Entry_repo {
	return &Entry_repo{repo: repo}
}
