package http

import (
	"errors"

	"github.com/uudashr/coursehub/internal/account"
)

// AccountService provides account related service.
type AccountService struct {
	repo account.Repository
}

// NewAccountService construtcs new AccountService.
func NewAccountService(repo account.Repository) (*AccountService, error) {
	if repo == nil {
		return nil, errors.New("nil repo")
	}

	return &AccountService{repo: repo}, nil
}

// RegisterNewAccount registers new account.
func (svc *AccountService) RegisterNewAccount(name, email string) (*account.Account, error) {
	id := account.NextID()
	acc, err := account.New(id, name, email, false)
	if err != nil {
		return nil, err
	}

	if err = svc.repo.Store(acc); err != nil {
		return nil, err
	}

	return acc, nil
}

// RetrieveAllAccounts listings all accounts.
func (svc *AccountService) RetrieveAllAccounts() ([]*account.Account, error) {
	return svc.repo.AllAccounts()
}

// RetrieveAccountDetails on the system.
func (svc *AccountService) RetrieveAccountDetails(id string) (*account.Account, error) {
	return svc.repo.AccountWithID(id)
}
