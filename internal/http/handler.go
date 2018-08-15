package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type delegate struct {
	accSvc *AccountService
}

func (d *delegate) registerNewAccount(c echo.Context) error {
	var payload newAccountPayload
	if err := c.Bind(&payload); err != nil {
		return err
	}

	acc, err := d.accSvc.RegisterNewAccount(payload.Name, payload.Email)
	if err != nil {
		return err
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/accounts/%s", acc.ID()))
	return c.NoContent(http.StatusCreated)
}

func (d *delegate) retrieveAllAccounts(c echo.Context) error {
	accs, err := d.accSvc.RetrieveAllAccounts()
	if err != nil {
		return err
	}

	respPayload := make([]accountPayload, len(accs))
	for i, acc := range accs {
		respPayload[i] = accountPayload{
			Name:     acc.Name(),
			Email:    acc.Email(),
			Verified: acc.Verified(),
		}
	}

	return c.JSON(http.StatusOK, respPayload)
}

func (d *delegate) ping(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// NewHandler constructs new http.Handler.
func NewHandler(accSvc *AccountService) (http.Handler, error) {
	if accSvc == nil {
		return nil, errors.New("nil accSvc")
	}
	d := &delegate{accSvc: accSvc}

	e := echo.New()
	e.POST("/accounts", d.registerNewAccount)
	e.GET("/accounts", d.retrieveAllAccounts)
	e.GET("/ping", d.ping)

	return e, nil
}

type newAccountPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type accountPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}
