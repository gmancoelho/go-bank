package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	m "github.com/gmancoelho/go-bank/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	account := &m.Account{
		FirstName: "John",
		LastName:  "Doe",
		Number:    12345,
		Balance:   1000,
	}

	// Mock the INSERT query
	mock.ExpectQuery(`INSERT INTO account`).
		WithArgs(account.FirstName, account.LastName, account.Number, account.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	// Call CreateAccount
	err = store.CreateAccount(account)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 1, account.ID)
	assert.NotZero(t, account.CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	// Mock the DELETE query
	mock.ExpectExec(`DELETE FROM account WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate 1 row affected

	// Call DeleteAccount
	err = store.DeleteAccount(1)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteAccountNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	// Mock the DELETE query with no rows affected
	mock.ExpectExec(`DELETE FROM account WHERE id = \$1`).
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(0, 0)) // Simulate 0 rows affected

	// Call DeleteAccount
	err = store.DeleteAccount(999)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "account id 999 not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAccounts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	// Mock the SELECT query
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "number", "balance", "created_at"}).
		AddRow(1, "John", "Doe", 12345, 1000, time.Now()).
		AddRow(2, "Jane", "Smith", 67890, 2000, time.Now())

	mock.ExpectQuery("(?i)select \\* from account").WillReturnRows(rows)

	// Call GetAccounts
	accounts, err := store.GetAccounts()

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, accounts, 2)
	assert.Equal(t, "John", accounts[0].FirstName)
	assert.Equal(t, "Jane", accounts[1].FirstName)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAccountByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	// Mock the SELECT query
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "number", "balance", "created_at"}).
		AddRow(1, "John", "Doe", 12345, 1000, time.Now())

	mock.ExpectQuery("(?i)select \\* from account where id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	// Call GetAccountByID
	account, err := store.GetAccountByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "John", account.FirstName)
	assert.Equal(t, int64(12345), account.Number)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAccountByIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	// Mock the SELECT query with no rows returned
	mock.ExpectQuery("(?i)select \\* from account where id = \\$1").
		WithArgs(999).
		WillReturnRows(sqlmock.NewRows(nil)) // No rows returned

	// Call GetAccountByID
	account, err := store.GetAccountByID(999)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, account)
	assert.EqualError(t, err, "account id 999 not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	account := &m.Account{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
	}

	// Mock the UPDATE query
	mock.ExpectExec("UPDATE account SET first_name = \\$1, last_name = \\$2 WHERE id = \\$3").
		WithArgs(account.FirstName, account.LastName, account.ID).
		WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate 1 row affected

	// Call UpdateAccount
	err = store.UpdateAccount(account)

	// Assertions
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateAccountNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	account := &m.Account{
		ID:        999,
		FirstName: "Jane",
		LastName:  "Doe",
	}

	// Mock the UPDATE query with no rows affected
	mock.ExpectExec("UPDATE account SET first_name = \\$1, last_name = \\$2 WHERE id = \\$3").
		WithArgs(account.FirstName, account.LastName, account.ID).
		WillReturnResult(sqlmock.NewResult(0, 0)) // Simulate 0 rows affected

	// Call UpdateAccount
	err = store.UpdateAccount(account)

	// Assertions
	assert.Error(t, err)
	assert.EqualError(t, err, "account id 999 not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateAccountQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	account := &m.Account{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
	}

	// Mock the UPDATE query with an error
	mock.ExpectExec("UPDATE account SET first_name = \\$1, last_name = \\$2 WHERE id = \\$3").
		WithArgs(account.FirstName, account.LastName, account.ID).
		WillReturnError(fmt.Errorf("query error"))

	// Call UpdateAccount
	err = store.UpdateAccount(account)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "query error")
	assert.NoError(t, mock.ExpectationsWereMet())
}
