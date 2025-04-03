package models

type TransferRequest struct {
	ToAccountID   int     `json: toAccountID`
	FromAccountID int     `json: fromAccountID`
	Amount        float64 `json: amout`
}
