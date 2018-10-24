package account_test

import (
	"fmt"
	"testing"

	"github.com/uudashr/coursehub/internal/account"
)

func TestAccount(t *testing.T) {
	cases := []struct {
		id       string
		name     string
		email    string
		verified bool
		valid    bool
	}{
		{id: "id-01", name: "Foo Bar", email: "foo.bar@kudo.co.id", verified: false, valid: true},
		{id: "id-02", name: "John Appleseed", email: "john.appleseed@kudo.co.id", verified: true, valid: true},
		{id: "", name: "John Appleseed", email: "john.appleseed@kudo.co.id", verified: true, valid: false},
		{id: "id-02", name: "", email: "john.appleseed@kudo.co.id", verified: true, valid: false},
		{id: "id-02", name: "John Appleseed", email: "", verified: true, valid: false},
		{id: "", name: "", email: "", verified: true, valid: false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			acc, err := account.New(c.id, c.name, c.email, c.verified)
			if c.valid {
				if err != nil {
					t.Fatal(err)
				}

				if got, want := acc.ID(), c.id; got != want {
					t.Errorf("id got: %q, want: %q", got, want)
				}

				if got, want := acc.Name(), c.name; got != want {
					t.Errorf("name got: %q, want: %q", got, want)
				}

				if got, want := acc.Email(), c.email; got != want {
					t.Errorf("email got: %q, want: %q", got, want)
				}

				if got, want := acc.Verified(), c.verified; got != want {
					t.Errorf("verified got: %v, want: %v", got, want)
				}
			} else {
				if err == nil {
					t.Error("expect error")
				}
			}
		})
	}
}

func TestNextID(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i <= 100; i++ {
		id := account.NextID()
		if _, found := m[id]; found {
			t.Fatalf("duplicate id %q on %v", id, m)
		}
	}
}
