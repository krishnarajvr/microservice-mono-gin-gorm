// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "micro/model"

	mock "github.com/stretchr/testify/mock"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// ListUsers provides a mock function with given fields:
func (_m *IUserRepo) ListUsers() (model.Users, error) {
	ret := _m.Called()

	var r0 model.Users
	if rf, ok := ret.Get(0).(func() model.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
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