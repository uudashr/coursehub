package account

import (
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
	return &Account{
		id:       id,
		name:     name,
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
type Repository interface {
	Store(*Account) error
	AllAccounts() ([]*Account, error)
}

// NextID returns next id.
func NextID() string {
	return bson.NewObjectId().Hex()
}
