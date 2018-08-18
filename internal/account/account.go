package account

import (
	"context"

	"github.com/uudashr/coursehub/eventkit"
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

// Create creates new account.
func Create(ctx context.Context, id, name, email string) (*Account, error) {
	acc, err := New(id, name, email, false)
	if err != nil {
		return nil, err
	}

	eventkit.PublishContext(ctx, AccountCreated{
		ID:    id,
		Name:  name,
		Email: email,
	})

	return acc, nil
}
