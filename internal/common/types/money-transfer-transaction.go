package types

type TransferMoneyTransactionRequest struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
}

type TransferMoneyTransactionResponse struct {
	ToAccountID   int64   `json:"to_account_id"`
	FromAccountID int64   `json:"from_account_id"`
	Amount        float64 `json:"currency"`
}
