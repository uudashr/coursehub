// +build integration

package mysql_test

import (
	"testing"

	"github.com/uudashr/coursehub/internal/account"
	"github.com/uudashr/coursehub/internal/mysql"
)

func TestAccountRepository_Store(t *testing.T) {
	fix := setupDBFixture(t)
	defer fix.tearDown()

	repo, err := mysql.NewAccountRepository(fix.db)
	if err != nil {
		t.Fatal(err)
	}

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
	fix := setupDBFixture(t)
	defer fix.tearDown()

	repo, err := mysql.NewAccountRepository(fix.db)
	if err != nil {
		t.Fatal(err)
	}

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

func TestAccountRepository_RetrieveAccount_notFound(t *testing.T) {
	fix := setupDBFixture(t)
	defer fix.tearDown()

	repo, err := mysql.NewAccountRepository(fix.db)
	if err != nil {
		t.Fatal(err)
	}

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
