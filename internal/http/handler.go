package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/uudashr/coursehub/internal/core"
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
		if err == core.ErrDuplicate {
			return c.NoContent(http.StatusConflict)
		}

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
			ID:       acc.ID(),
			Name:     acc.Name(),
			Email:    acc.Email(),
			Verified: acc.Verified(),
		}
	}

	return c.JSON(http.StatusOK, respPayload)
}

func (d *delegate) retrieveAccountDetails(c echo.Context) error {
	paramID := c.Param("id")
	acc, err := d.accSvc.RetrieveAccountDetails(paramID)
	if err != nil {
		return err
	}

	if acc == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, accountPayload{
		ID:       acc.ID(),
		Name:     acc.Name(),
		Email:    acc.Email(),
		Verified: acc.Verified(),
	})
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
	e.GET("/accounts/:id", d.retrieveAccountDetails)
	e.GET("/ping", d.ping)

	return e, nil
}

type newAccountPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type accountPayload struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
}
