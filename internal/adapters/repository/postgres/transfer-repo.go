package postgres

// import (
// 	"gobanking/internal/core/domain"
// 	"log"
// )

// type Transfer_repo struct {
// 	repo *Postgres_repo
// }

// func (tr *Transfer_repo) Create(transfer *domain.Transfer) error {
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	query := `
// 		INSERT INTO transfers (to_account, from_account, amount)
// 		VALUES ($1, $2, $3)
// 	`
// 	_, err := tr.repo.db.ExecContext(ctx, query, transfer.ToAccountID, transfer.FromAccountID, transfer.Amount)

// 	if err != nil {
// 		log.Printf("error during the execution of the INSERT query at the repository layer : %v \n", err)
// 		return err
// 	}

// 	return nil
// }

// func (tr *Transfer_repo) GetAll() ([]*domain.Transfer, error) {
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	query := `SELECT * FROM transfers`

// 	rows, err := tr.repo.db.QueryContext(ctx, query)
// 	if err != nil {
// 		log.Printf("error during the execution of the SELECT * query at the repository layer : %v \n", err)
// 		return nil, err
// 	}

// 	var transfers []*domain.Transfer

// 	for rows.Next() {
// 		var transfer domain.Transfer
// 		err := rows.Scan(
// 			&transfer.ID,
// 			&transfer.ToAccountID,
// 			&transfer.FromAccountID,
// 			&transfer.Amount,
// 			&transfer.CreatedAt,
// 			&transfer.UpdatedAt,
// 		)
// 		if err != nil {
// 			log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
// 			return nil, err
// 		}

// 		transfers = append(transfers, &transfer)
// 	}

// 	return transfers, nil
// }

// func (tr *Transfer_repo) GetById(id int) (*domain.Transfer, error) {

// 	// create ctx
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	// define the query
// 	query := `
// 		SELECT * FROM transfers
// 		WHERE id = $1
// 	`

// 	row := tr.repo.db.QueryRowContext(ctx, query, id)

// 	var transfer domain.Transfer
// 	err := row.Scan(
// 		&transfer.ID,
// 		&transfer.ToAccountID,
// 		&transfer.FromAccountID,
// 		&transfer.Amount,
// 		&transfer.CreatedAt,
// 		&transfer.UpdatedAt,
// 	)

// 	if err != nil {
// 		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
// 		return nil, err
// 	}

// 	return &transfer, nil
// }

// func (tr *Transfer_repo) GetByToAccountId(to_account_id int) (*domain.Transfer, error) {
// 	// create ctx
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	// define the query
// 	query := `
// 		SELECT * FROM transfers
// 		WHERE to_account = $1
// 	`

// 	row := tr.repo.db.QueryRowContext(ctx, query, to_account_id)

// 	var transfer domain.Transfer
// 	err := row.Scan(
// 		&transfer.ID,
// 		&transfer.ToAccountID,
// 		&transfer.FromAccountID,
// 		&transfer.Amount,
// 		&transfer.CreatedAt,
// 		&transfer.UpdatedAt,
// 	)

// 	if err != nil {
// 		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
// 		return nil, err
// 	}

// 	return &transfer, nil
// }

// func (tr *Transfer_repo) GetByFromAccountId(from_account_id int) (*domain.Transfer, error) {
// 	// create ctx
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	// define the query
// 	query := `
// 		SELECT * FROM transfers
// 		WHERE from_account = $1
// 	`

// 	row := tr.repo.db.QueryRowContext(ctx, query, from_account_id)

// 	var transfer domain.Transfer
// 	err := row.Scan(
// 		&transfer.ID,
// 		&transfer.ToAccountID,
// 		&transfer.FromAccountID,
// 		&transfer.Amount,
// 		&transfer.CreatedAt,
// 		&transfer.UpdatedAt,
// 	)

// 	if err != nil {
// 		log.Printf("error during the mapping between the domain entity and db table at the repository layer : %v \n", err)
// 		return nil, err
// 	}

// 	return &transfer, nil
// }

// func (tr *Transfer_repo) GetByFromAndToAccountId(from_acc_id, to_acc_id int) (*domain.Transfer, error) {
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	query := `SELECT * FROM transfers WHERE to_account = $1 AND from_account=$2`

// 	row := tr.repo.db.QueryRowContext(ctx, query, to_acc_id, from_acc_id)
// 	var transfer_record domain.Transfer
// 	err := row.Scan(
// 		&transfer_record.ID,
// 		transfer_record.ToAccountID,
// 		transfer_record.FromAccountID,
// 		transfer_record.Amount,
// 		transfer_record.CreatedAt,
// 		transfer_record.UpdatedAt,
// 	)
// 	if err != nil {
// 		log.Println("error while fetching transfer record from database based on to and from accounts ids : ", err)
// 		return nil, err
// 	}
// 	return &transfer_record, nil
// }

// func (tr *Transfer_repo) Delete(id int) error {
// 	ctx, cancel := CreateContext()
// 	defer cancel()

// 	query := `DELETE FROM transfers WHERE id = $1`

// 	_, err := tr.repo.db.ExecContext(ctx, query, id)

// 	if err != nil {
// 		log.Printf("error during the execution of the DELETE query inside postgres repository layer : %v \n", err)
// 		return err
// 	}

// 	return nil

// }

// func NewTransferRepo(repo *Postgres_repo) *Transfer_repo {
// 	return &Transfer_repo{repo: repo}
// }
