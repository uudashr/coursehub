package mysql

import (
	"database/sql"
	"errors"

	"github.com/uudashr/coursehub/internal/account"
)

// AccountRepository is repository for Account.
type AccountRepository struct {
	db *sql.DB
}

// NewAccountRepository creates new AccountRepository.
func NewAccountRepository(db *sql.DB) (*AccountRepository, error) {
	if db == nil {
		return nil, errors.New("")
	}

	return &AccountRepository{db}, nil
}

// Store acc to repository.
func (r *AccountRepository) Store(acc *account.Account) error {
	res, err := r.db.Exec("INSERT INTO accounts (id, name, email, verified) VALUES (?, ?, ?, ?)", acc.ID(), acc.Email(), acc.Name(), acc.Verified())
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("No rows affected")
	}

	return nil
}

// AllAccounts on the repository.
func (r *AccountRepository) AllAccounts() ([]*account.Account, error) {
	rows, err := r.db.Query("SELECT id, name, email, verified FROM accounts")
	if err != nil {
		return nil, err
	}

	var accounts []*account.Account
	for rows.Next() {
		var (
			id       string
			name     string
			email    string
			verified bool
		)
		if err = rows.Scan(
			&id,
			&name,
			&email,
			&verified,
		); err != nil {
			return nil, err
		}

		acc, err := account.New(id, name, email, verified)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

// AccountWithID implements the account.Repository interface.
func (r *AccountRepository) AccountWithID(id string) (*account.Account, error) {
	var (
		name     string
		email    string
		verified bool
	)

	err := r.db.QueryRow("SELECT name, email, verified FROM accounts WHERE id = ?", id).Scan(
		&name,
		&email,
		&verified,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return account.New(id, name, email, verified)
}
