// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/uudashr/coursehub/internal/account"
import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AccountWithID provides a mock function with given fields: id
func (_m *Repository) AccountWithID(id string) (*account.Account, error) {
	ret := _m.Called(id)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(string) *account.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllAccounts provides a mock function with given fields:
func (_m *Repository) AllAccounts() ([]*account.Account, error) {
	ret := _m.Called()

	var r0 []*account.Account
	if rf, ok := ret.Get(0).(func() []*account.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0
func (_m *Repository) Store(_a0 *account.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}