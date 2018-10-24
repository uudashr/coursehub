package inmem_test

import (
	"testing"

	"github.com/uudashr/coursehub/internal/account"
	"github.com/uudashr/coursehub/internal/inmem"
)

func TestAccountRepository_Store(t *testing.T) {
	repo := inmem.NewAccountRepository()

	acc, err := account.New("some-id", "Umar", "umar@kudo.co.id", false)
	if err != nil {
		t.Fatal(err)
	}

	prevAccs, err := repo.AllAccounts()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Store(acc)
	if err != nil {
		t.Fatal(err)
	}

	accs, err := repo.AllAccounts()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(accs), len(prevAccs)+1; got != want {
		t.Error("len got:", got, "want:", want)
	}
}

func TestAccountRepository_Store_duplicates(t *testing.T) {
	repo := inmem.NewAccountRepository()

	acc, err := account.New("some-id", "Umar", "umar@kudo.co.id", false)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Store(acc)
	if err != nil {
		t.Fatal(err)
	}

	prevAccs, err := repo.AllAccounts()
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Store(acc)
	if err == nil {
		t.Fatal("should fail on duplicate id")
	}

	accs, err := repo.AllAccounts()
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(accs), len(prevAccs); got != want {
		t.Error("len got:", got, "want:", want)
	}
}

func TestAccountRepository_RetrieveStoredAccount(t *testing.T) {
	repo := inmem.NewAccountRepository()

	acc, err := account.New("some-id", "Umar", "umar@kudo.co.id", false)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Store(acc)
	if err != nil {
		t.Fatal(err)
	}

	retAcc, err := repo.AccountWithID(acc.ID())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := retAcc.ID(), acc.ID(); got != want {
		t.Errorf("id got: %q, want: %q", got, want)
	}

	if got, want := retAcc.Name(), acc.Name(); got != want {
		t.Errorf("name got: %q, want: %q", got, want)
	}

	if got, want := retAcc.Email(), acc.Email(); got != want {
		t.Errorf("email got: %q, want: %q", got, want)
	}

	if got, want := retAcc.Verified(), acc.Verified(); got != want {
		t.Errorf("verified got: %t, want: %t", got, want)
	}
}

func TestAccountRepository_RetrieveAccount_notFound(t *testing.T) {
	repo := inmem.NewAccountRepository()

	acc, err := account.New("some-id", "Umar", "umar@kudo.co.id", false)
	if err != nil {
		t.Fatal(err)
	}

	retAcc, err := repo.AccountWithID(acc.ID())
	if err != nil {
		t.Fatal(err)
	}

	if retAcc != nil {
		t.Fatal("expect not found")
	}
}
