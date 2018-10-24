package fixture

import (
	"github.com/icrowley/fake"
	"github.com/uudashr/coursehub/internal/account"
)

// Accounts fixture.
func Accounts(n int) []*account.Account {
	accs := make([]*account.Account, n)
	for i := 0; i < n; i++ {
		accs[i] = Account()
	}
	return accs
}

// Account fixture.
func Account() *account.Account {
	acc, err := account.New(account.NextID(), fake.FullName(), fake.EmailAddress(), false)
	if err != nil {
		panic(err)
	}

	return acc
}
