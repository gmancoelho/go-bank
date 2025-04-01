package repository

import (
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

	mock.ExpectQuery(`INSERT INTO account`).
		WithArgs(account.FirstName, account.LastName, account.Number, account.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	err = store.CreateAccount(account)
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

	mock.ExpectExec(`DELETE FROM account WHERE id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = store.DeleteAccount(1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAccounts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	store := &PostgressStore{db: db}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "number", "balance", "created_at"}).
		AddRow(1, "John", "Doe", 12345, 1000, time.Now()).
		AddRow(2, "Jane", "Smith", 67890, 2000, time.Now())

	mock.ExpectQuery("(?i)select \\* from account").WillReturnRows(rows)

	accounts, err := store.GetAccounts()
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

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "number", "balance", "created_at"}).
		AddRow(1, "John", "Doe", int(12345), 1000, time.Now())

	mock.ExpectQuery("(?i)select \\* from account where id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	account, err := store.GetAccountByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "John", account.FirstName)
    assert.Equal(t, int64(12345), account.Number)

	assert.NoError(t, mock.ExpectationsWereMet())
}
