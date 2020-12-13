// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "micro/model"

	mock "github.com/stretchr/testify/mock"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *IUserService) GetAll() (model.UserDtos, error) {
	ret := _m.Called()

	var r0 model.UserDtos
	if rf, ok := ret.Get(0).(func() model.UserDtos); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UserDtos)
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
