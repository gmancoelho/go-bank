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
	GetAccountByID(int) (*m.Account, error)
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

func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgressStore) createAccountTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS account (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        number BIGINT NOT NULL UNIQUE,
        balance BIGINT NOT NULL DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgressStore) CreateAccount(*m.Account) error {
	return nil
}

func (s *PostgressStore) DeleteAccount(int) error {
	return nil

}

func (s *PostgressStore) UpdateAccount(*m.Account) error {
	return nil

}

func (s *PostgressStore) GetAccountByID(int) (*m.Account, error) {
	return nil, nil
}
