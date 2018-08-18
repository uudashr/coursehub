package http

import (
	"context"
	"errors"

	"github.com/uudashr/coursehub/applifecycle"

	"github.com/uudashr/coursehub/internal/account"
)

// AccountService provides account related service.
type AccountService struct {
	repo          account.Repository
	EventsHandler applifecycle.EventsHandler
}

// NewAccountService construtcs new AccountService.
func NewAccountService(repo account.Repository) (*AccountService, error) {
	if repo == nil {
		return nil, errors.New("nil repo")
	}

	return &AccountService{repo: repo}, nil
}

// RegisterNewAccount registers new account.
func (svc *AccountService) RegisterNewAccount(name, email string) (acc *account.Account, err error) {
	lc := &applifecycle.Lifecycle{
		Handler: svc.EventsHandler,
	}
	defer lc.End(err)

	id := account.NextID()
	acc, err = account.Create(lc.Context(context.Background()), id, name, email)
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
