package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	nethttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/uudashr/coursehub/internal/account"
	accountmocks "github.com/uudashr/coursehub/internal/account/mocks"
	"github.com/uudashr/coursehub/internal/core"
	modelfixture "github.com/uudashr/coursehub/internal/fixture"
	"github.com/uudashr/coursehub/internal/http"
)

type fixture struct {
	t       *testing.T
	handler nethttp.Handler
	repo    *accountmocks.Repository
}

func setupFixture(t *testing.T) *fixture {
	repo := new(accountmocks.Repository)
	svc, err := http.NewAccountService(repo)
	if err != nil {
		panic(err)
	}

	handler, err := http.NewHandler(svc)
	if err != nil {
		panic(err)
	}

	return &fixture{
		t:       t,
		repo:    repo,
		handler: handler,
	}
}

func (fix *fixture) tearDown() {
	fix.repo.AssertExpectations(fix.t)
}

func TestNewHandler(t *testing.T) {
	_, err := http.NewHandler(nil)
	if err == nil {
		t.Fatal("expect error")
	}
}

func TestHandler_Ping(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	res := httpGet(fix.handler, "/ping")
	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}
}

func TestHandler_RegisterNewAccount(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	name := "Nuruddin Ashr"
	email := "nuruddin.ashr@kudo.co.id"

	fix.repo.On("Store", mock.MatchedBy(func(acc *account.Account) bool {
		return acc.Name() == name && acc.Email() == email
	})).Return(nil)

	res := httpPost(fix.handler, "/accounts", map[string]interface{}{
		"name":  name,
		"email": email,
	})

	if got, want := res.StatusCode, nethttp.StatusCreated; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}
}

func TestHandler_RegisterNewAccount_conflict(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	fix.repo.On("Store", mock.Anything).Return(core.ErrDuplicate)

	name := "Nuruddin Ashr"
	email := "nuruddin.ashr@kudo.co.id"
	res := httpPost(fix.handler, "/accounts", map[string]interface{}{
		"name":  name,
		"email": email,
	})

	if got, want := res.StatusCode, nethttp.StatusConflict; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}
}

func TestHandler_RetrieveAllAccounts(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	storedAccs := modelfixture.Accounts(10)
	fix.repo.On("AllAccounts").Return(storedAccs, nil)

	res := httpGet(fix.handler, "/accounts")

	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}

	var accs []accountPayload
	if err := json.NewDecoder(res.Body).Decode(&accs); err != nil {
		t.Fatal(err)
	}

	if got, want := len(accs), len(storedAccs); got != want {
		t.Fatalf("len got: %d, want: %d", got, want)
	}

	for i := range accs {
		storedAcc := storedAccs[i]
		acc := accs[i]

		if got, want := storedAcc.ID(), acc.ID; got != want {
			t.Errorf("id got: %q, want: %q, index: %d", got, want, i)
		}

		if got, want := storedAcc.Name(), acc.Name; got != want {
			t.Errorf("name got: %q, want: %q, index: %d", got, want, i)
		}

		if got, want := storedAcc.Email(), acc.Email; got != want {
			t.Errorf("email got: %q, want: %q, index: %d", got, want, i)
		}

		if got, want := storedAcc.Verified(), acc.Verified; got != want {
			t.Errorf("verified got: %t, want: %t, index: %d", got, want, i)
		}
	}
}

func TestHandler_RetrieveAccountDetails(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	storedAcc := modelfixture.Account()
	fix.repo.On("AccountWithID", storedAcc.ID()).Return(storedAcc, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/accounts/%s", storedAcc.ID()))

	if got, want := res.StatusCode, nethttp.StatusOK; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}

	var acc accountPayload
	if err := json.NewDecoder(res.Body).Decode(&acc); err != nil {
		t.Fatal(err)
	}

	if got, want := storedAcc.ID(), acc.ID; got != want {
		t.Errorf("id got: %q, want: %q", got, want)
	}

	if got, want := storedAcc.Name(), acc.Name; got != want {
		t.Errorf("name got: %q, want: %q", got, want)
	}

	if got, want := storedAcc.Email(), acc.Email; got != want {
		t.Errorf("email got: %q, want: %q", got, want)
	}

	if got, want := storedAcc.Verified(), acc.Verified; got != want {
		t.Errorf("verified got: %t, want: %t", got, want)
	}
}

func TestHandler_RetrieveAccountDetails_notFound(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	id := account.NextID()
	fix.repo.On("AccountWithID", id).Return(nil, nil)

	res := httpGet(fix.handler, fmt.Sprintf("/accounts/%s", id))

	if got, want := res.StatusCode, nethttp.StatusNotFound; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}
}

type accountPayload struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}

func httpGet(h nethttp.Handler, path string) *nethttp.Response {
	req := httptest.NewRequest(nethttp.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Result()
}

func httpPost(h nethttp.Handler, path string, body interface{}) *nethttp.Response {
	bodyReader, err := makeReader(body)
	if err != nil {
		panic(err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(nethttp.MethodPost, path, bodyReader)
	if bodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	h.ServeHTTP(rec, req)
	return rec.Result()
}

// convert body to json as reader
func makeReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}

	switch v := body.(type) {
	case string:
		// return as is
		return strings.NewReader(v), nil
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		return bytes.NewReader(b), nil
	}
}
