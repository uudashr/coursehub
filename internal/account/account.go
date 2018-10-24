package account

import (
	"errors"

	"labix.org/v2/mgo/bson"
)

// Account represents user/organization account.
type Account struct {
	id       string
	name     string
	email    string
	verified bool
}

// New creates new Account instance.
func New(id, name, email string, verified bool) (*Account, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if name == "" {
		return nil, errors.New("empty name")
	}

	if email == "" {
		return nil, errors.New("empty email")
	}

	return &Account{
		id:       id,
		name:     name,
		email:    email,
		verified: verified,
	}, nil
}

// ID of the account.
func (acc Account) ID() string {
	return acc.id
}

// Name of the account.
func (acc Account) Name() string {
	return acc.name
}

// Email of the account.
func (acc Account) Email() string {
	return acc.email
}

// Verified status of the account.
func (acc Account) Verified() bool {
	return acc.verified
}

// Repository represents Account repository.
//go:generate mockery -name=Repository
type Repository interface {
	Store(*Account) error
	AllAccounts() ([]*Account, error)
	AccountWithID(id string) (*Account, error)
}

// NextID returns next id.
func NextID() string {
	return bson.NewObjectId().Hex()
}
