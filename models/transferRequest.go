package models

type TransferRequest struct {
	ToAccount   int     `json: toAccount`
	FromAccount int     `json: fromAccount`
	Amount      float64 `json: amout`
}
