package postgres

import (
	"fmt"
	"gobanking/internal/core/domain"
	"log"
	"strings"
)

type Account_repo struct {
	repo *Postgres_repo
}

// CREATE REQUEST
func (ar *Account_repo) Create(account domain.Account) error {
	ctx, cancel := CreateContext()
	defer cancel()

	// define the query
	query := `
		INSERT INTO accounts (owner_name, balance, currency)
		VALUES 
		($1, $2, $3)
	`

	// execute the query and check for errors
	_, err := ar.repo.db.ExecContext(ctx, query, ar, account.OwnerName, account.Balance, account.Currency)
	if err != nil {
		log.Printf("error during the execution of the INSERT query at the repository layer : %v \n", err)
		return err
	}

	return nil
}

// GET ALL REQUEST
func (ar *Account_repo) GetAll() ([]*domain.Account, error) {

	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define the query
	query := `
		SELECT * FROM accounts
	`

	rows, err := ar.repo.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("error during the execution of the SELECT * query at the repository layer : %v \n", err)
		return nil, err
	}

	var accounts []*domain.Account

	for rows.Next() {
		var account domain.Account
		err := rows.Scan(
			&account.ID,
			&account.OwnerName,
			&account.Balance,
			&account.Currency,
			&account.Deleted,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
			return nil, err
		}

		accounts = append(accounts, &account)
	}

	return accounts, nil
}

// GET BY ID REQUEST
func (ar *Account_repo) GetById(id int) (*domain.Account, error) {

	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define the query
	query := `
		SELECT * FROM accounts
		WHERE id = $1
	`

	row := ar.repo.db.QueryRowContext(ctx, query, id)

	var account domain.Account
	err := row.Scan(&account.ID,
		&account.OwnerName,
		&account.Balance,
		&account.Currency,
		&account.Deleted,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
		return nil, err
	}

	return &account, nil
}

// GET BY ID REQUEST
func (ar *Account_repo) GetByOwnerName(ownerName string) (*domain.Account, error) {

	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define the query
	query := `
		SELECT * FROM accounts
		WHERE owner_name = $1
	`

	row := ar.repo.db.QueryRowContext(ctx, query, ownerName)

	var account domain.Account
	err := row.Scan(&account.ID,
		&account.OwnerName,
		&account.Balance,
		&account.Currency,
		&account.Deleted,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
		return nil, err
	}

	return &account, nil
}

// UPDATE REQUEST
// => only the balance and the deleted fields are allowed to be modified, but we can't modify the account owner or the account currency type
type Update_account_params struct {
	Balance  *float64 // i made these types as * to the go-type because i need to check if these are provided or not
	Currency *string
}

func (ar *Account_repo) Update(id int, update_params Update_account_params) (*domain.Account, error) {
	// create ctx
	ctx, cancel := CreateContext()
	defer cancel()

	// define the query
	var (
		updateQuery strings.Builder
		params      []interface{}
		paramCount  int = 1
	)

	updateQuery.WriteString("UPDATE accounts SET ")
	// validation that the balance is there and its value is not less than 0
	if update_params.Balance != nil && !(*update_params.Balance < 0) {
		updateQuery.WriteString(fmt.Sprintf("balance = $%d,", paramCount))
		paramCount++
		// append the value to pass it to the query by the order of the count it took in the line before
		params = append(params, *update_params.Balance)
	}
	// validation that the currency is there and not empty
	if update_params.Currency != nil && !(*update_params.Currency == "") {
		updateQuery.WriteString(fmt.Sprintf("currency = $%d ", paramCount))
		paramCount++
		// append the value to pass it to the query by the order of the count it took in the line before
		params = append(params, *update_params.Currency)
	}

	updateQuery.WriteString(fmt.Sprintf("WHERE id = $%d", paramCount))
	params = append(params, id)

	_, err := ar.repo.db.ExecContext(ctx, updateQuery.String(), params...)
	if err != nil {
		log.Printf("error during the execution of the UPDATE query from the postgres repository layer : %v \n ", err)
		return nil, err
	}

	// fetch back the updated acc from database
	updated_account, _ := ar.GetById(id)

	return updated_account, nil
}

// DELETE REQUEST
func (ar *Account_repo) Delete(id int) error {
	ctx, cancel := CreateContext()
	defer cancel()

	query := `DELETE FROM accounts WHERE id = $1`

	_, err := ar.repo.db.ExecContext(ctx, query, id)

	if err != nil {
		log.Printf("error during the execution of the DELETE query inside postgres repository layer : %v \n", err)
		return err
	}

	return nil
}

func NewAccountRepo(repo *Postgres_repo) *Account_repo {
	return &Account_repo{repo: repo}
}
