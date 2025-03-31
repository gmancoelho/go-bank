package api

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func newAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
	}
}

func (a *Account) GetID() int {
	return a.ID
}

func (a *Account) GetName() string {
	return a.FirstName + " " + a.LastName
}
