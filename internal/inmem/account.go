package inmem

import (
	"errors"

	"github.com/uudashr/coursehub/internal/account"
)

// AccountRepository represents the Account repository.
type AccountRepository struct {
	m map[string]account.Account
}

// NewAccountRepository constructs new AccountRepository.
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{m: make(map[string]account.Account)}
}

// Store implements the account.Repository interface.
func (r *AccountRepository) Store(acc *account.Account) error {
	if _, found := r.m[acc.ID()]; found {
		return errors.New("duplicate id")
	}

	r.m[acc.ID()] = *acc
	return nil
}

// AllAccounts implements the account.Repository interface.
func (r *AccountRepository) AllAccounts() ([]*account.Account, error) {
	accs := make([]*account.Account, len(r.m))
	var i int
	for _, v := range r.m {
		acc := v
		accs[i] = &acc
		i++
	}

	return accs, nil
}
