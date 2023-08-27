package account

import (
	"context"
	"fmt"
	"log"

	"github.com/FadyGamilM/go-banking-v2/db"
	domain "github.com/FadyGamilM/go-banking-v2/internal/account/domain"
)

type accountRepo struct {
	pg *db.PG
}

func NewAccountRepo(pg *db.PG) domain.AccountRepo {
	return &accountRepo{
		pg: pg,
	}
}

const (
	create_Query = `
		INSERT INTO accounts (owner_name, balance, currency)
		VALUES 
		($1, $2, $3)
		RETURNING id, owner_name, balance, currency, activated, created_at, updated_at
	`

	get_All_Query = `
		SELECT id, owner_name, balance, currency, activated, created_at, updated_at
		FROM accounts 
	`

	get_By_ID_Query = `
		SELECT id, owner_name, balance, currency, activated, created_at, updated_at
		FROM accounts
		WHERE id = $1
		FOR NO KEY UPDATE
	`

	get_By_OwnerName_Query = `
		SELECT id, owner_name, balance, currency, activated, created_at, updated_at
		FROM accounts
		WHERE owner_name = $1
	`

	delete_By_ID_Query = `
		DELETE FROM accounts 
		WHERE id = $1
	`

	delete_By_OwnerName_Query = `
		DELETE FROM accounts 
		WHERE owner_name = $1
	`

	update_Balance_By_id_Query = `
		UPDATE accounts 
		SET balance = balance + $1 
		WHERE id = $2
		RETURNING id, owner_name, balance, currency, activated, created_at, updated_at 
	`

	get_Account_Balance_Query = `
		SELECT balance
		FROM accounts 
		WHERE id = $1
	`
)

// Create new Account
func (repo *accountRepo) Create(ctx context.Context, acc *domain.Account) (*domain.Account, error) {
	createdAccount := new(domain.Account)
	// execute query and scan result
	err := repo.pg.DB.QueryRowContext(ctx, create_Query, acc.OwnerName, acc.Balance, acc.Currency).Scan(
		&createdAccount.ID,
		&createdAccount.OwnerName,
		&createdAccount.Balance,
		&createdAccount.Currency,
		&createdAccount.Activated,
		&createdAccount.CreatedAt,
		&createdAccount.UpdatedAt,
	)
	if err != nil {
		return &domain.Account{}, err
	}

	return createdAccount, nil
}

// get all accounts
func (repo *accountRepo) GetAll(ctx context.Context) ([]*domain.Account, error) {
	// execute query
	rows, err := repo.pg.DB.QueryContext(ctx, get_All_Query)
	if err != nil {
		return nil, err
	}

	// define domain entities to return after mapping the db entity to domain entity
	accounts := []*domain.Account{}
	for rows.Next() {
		// define db entity to scan the query result
		account := new(domain.Account)
		err := rows.Scan(
			&account.ID,
			&account.OwnerName,
			&account.Balance,
			&account.Currency,
			&account.Activated,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return []*domain.Account{}, err
		}
		// map and append the result
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (repo *accountRepo) GetByID(ctx context.Context, id int64) (*domain.Account, error) {
	// define database entity type to scan the query result
	account := new(domain.Account)

	log.Println("the id in the repo layer is : ", id)
	err := repo.pg.DB.QueryRowContext(ctx, get_By_ID_Query, id).Scan(
		&account.ID,
		&account.OwnerName,
		&account.Balance,
		&account.Currency,
		&account.Activated,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repo *accountRepo) GetByOwnerName(ctx context.Context, owner string) (*domain.Account, error) {
	// define database entity type to scan the query result
	account := new(domain.Account)
	err := repo.pg.DB.QueryRowContext(ctx, get_By_OwnerName_Query, owner).Scan(
		&account.ID,
		&account.OwnerName,
		&account.Balance,
		&account.Currency,
		&account.Activated,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		return &domain.Account{}, err
	}

	return account, nil
}

func (repo *accountRepo) DeleteByID(ctx context.Context, id int64) error {
	_, err := repo.pg.DB.ExecContext(ctx, delete_By_ID_Query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *accountRepo) UpdateByID(ctx context.Context, id int64, updateAmount float64) (*domain.Account, error) {
	account := new(domain.Account)

	// In case that the updateAmount is negative (withdraw operation) we need to validate that the account has at least a balance >= the updateAmount
	var current_balance float64
	err := repo.pg.DB.QueryRowContext(ctx, get_Account_Balance_Query, id).Scan(&current_balance)
	if err != nil {
		return nil, err
	}

	if updateAmount < 0 {
		if current_balance <= -updateAmount {
			return nil, fmt.Errorf("account balance is less than the amount you want to withdraw, balance is : %v ", current_balance)
		}
	}

	// update the balance to be equale the old balance + (-/+ Amount)
	err = repo.pg.DB.QueryRowContext(ctx, update_Balance_By_id_Query, updateAmount, id).Scan(&account.ID, &account.OwnerName, &account.Balance, &account.Currency, &account.Activated, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return account, nil
}
