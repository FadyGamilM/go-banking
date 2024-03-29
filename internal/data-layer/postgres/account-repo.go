package postgres

import (
	"fmt"
	"gobanking/internal/core-layer/domain"
	"gobanking/internal/data-layer/postgres/models"
	"gobanking/internal/infra-layer/db/postgres"
	"log"
)

type PG_AccountRepository struct {
	pg_tx *postgres.PG_TX
}

func NewPG_AccountRepo(tx *postgres.PG_TX) *PG_AccountRepository {
	return &PG_AccountRepository{
		pg_tx: tx,
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
		SET balance = $1 
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
func (repo *PG_AccountRepository) Create(acc *domain.Account) (*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define database entity type to scan the query result
	acc_db := new(models.PgAccount)

	// execute query and scan result
	err := repo.pg_tx.TX.QueryRowContext(ctx, create_Query, acc.OwnerName, acc.Balance, acc.Currency).Scan(&acc_db.ID, &acc_db.OwnerName, &acc_db.Balance, &acc_db.Currency, &acc_db.Activated, &acc_db.CreatedAt, &acc_db.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// modify the domain entity type to send it back to the service layer
	acc = acc_db.ToDomainEntity()

	return acc, nil
}

// get all accounts
func (repo *PG_AccountRepository) GetAll() ([]*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// execute query
	rows, err := repo.pg_tx.TX.QueryContext(ctx, get_All_Query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// define domain entities to return after mapping the db entity to domain entity
	domain_accounts := []*domain.Account{}
	for rows.Next() {
		// define db entity to scan the query result
		db_account := new(models.PgAccount)
		err := rows.Scan(
			&db_account.ID,
			&db_account.OwnerName,
			&db_account.Balance,
			&db_account.Currency,
			&db_account.Activated,
			&db_account.CreatedAt,
			&db_account.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// map and append the result
		domain_accounts = append(domain_accounts, db_account.ToDomainEntity())
	}

	return domain_accounts, nil
}

func (repo *PG_AccountRepository) GetByID(id int64) (*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define database entity type to scan the query result
	db_account := new(models.PgAccount)
	err := repo.pg_tx.TX.QueryRowContext(ctx, get_By_ID_Query, id).Scan(
		&db_account.ID,
		&db_account.OwnerName,
		&db_account.Balance,
		&db_account.Currency,
		&db_account.Activated,
		&db_account.CreatedAt,
		&db_account.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// define a domain entity type to return to service layer
	domain_account := db_account.ToDomainEntity()
	log.Println("repo will return this result => ", domain_account.OwnerName)
	return domain_account, nil
}

func (repo *PG_AccountRepository) GetByOwnerName(owner string) (*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define database entity type to scan the query result
	db_account := new(models.PgAccount)
	err := repo.pg_tx.TX.QueryRowContext(ctx, get_By_OwnerName_Query, owner).Scan(
		&db_account.ID,
		&db_account.OwnerName,
		&db_account.Balance,
		&db_account.Currency,
		&db_account.Activated,
		&db_account.CreatedAt,
		&db_account.UpdatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// define a domain entity type to return to service layer
	domain_account := db_account.ToDomainEntity()
	return domain_account, nil
}

func (repo *PG_AccountRepository) DeleteByID(id int64) error {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	_, err := repo.pg_tx.TX.ExecContext(ctx, delete_By_ID_Query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo *PG_AccountRepository) UpdateByID(id int64, updateAmount float64) (*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	db_acc := new(models.PgAccount)

	// In case that the updateAmount is negative (withdraw operation) we need to validate that the account has at least a balance >= the updateAmount
	var current_balance float64
	err := repo.pg_tx.TX.QueryRowContext(ctx, get_Account_Balance_Query, id).Scan(&current_balance)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("amoutn to be withdrawn or deposited is : ", updateAmount)

	if updateAmount < 0 {
		log.Println("yeeeeee 1")
		log.Println(current_balance)
		log.Println(-1 * updateAmount)
		if current_balance <= updateAmount*-1 {
			log.Println("yeeeeee 2")
			return nil, fmt.Errorf("account balance is less than the amount you want to withdraw, balance is : %v ", current_balance)
		}
	}

	// update the balance to be equale the old balance + (-/+ Amount)
	err = repo.pg_tx.TX.QueryRowContext(ctx, update_Balance_By_id_Query, updateAmount+current_balance, id).Scan(&db_acc.ID, &db_acc.OwnerName, &db_acc.Balance, &db_acc.Currency, &db_acc.Activated, &db_acc.CreatedAt, &db_acc.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	domain_acc := db_acc.ToDomainEntity()

	return domain_acc, nil
}
