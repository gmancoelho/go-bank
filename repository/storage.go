package repository

import (
	"database/sql"
	"fmt"

	m "github.com/gmancoelho/go-bank/models"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*m.Account) error
	DeleteAccount(int) error
	UpdateAccount(*m.Account) error
	GetAccountByID(int) (m.Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgressStore, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", DBUser, DBName, DBPassword, DBSSLMode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{db: db}, nil
}
