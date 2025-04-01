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
	GetAccounts() ([]*m.Account, error)
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

func (s *PostgressStore) CreateAccount(account *m.Account) error {
	query := `
        INSERT INTO account (first_name, last_name, number, balance) 
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at;
    `

	err := s.db.QueryRow(query, account.FirstName, account.LastName, account.Number, account.Balance).
		Scan(&account.ID, &account.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}

func (s *PostgressStore) DeleteAccount(int) error {
	return nil

}

func (s *PostgressStore) UpdateAccount(*m.Account) error {
	return nil

}

func (s *PostgressStore) GetAccounts() ([]*m.Account, error) {
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := make([]*m.Account, 0)

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	defer rows.Close()

	return accounts, nil
}

func (s *PostgressStore) GetAccountByID(id int) (*m.Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		return account, err
	}

	return nil, fmt.Errorf("account id %d not found", id)
}

func scanIntoAccount(rows *sql.Rows) (*m.Account, error) {
	account := new(m.Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}
