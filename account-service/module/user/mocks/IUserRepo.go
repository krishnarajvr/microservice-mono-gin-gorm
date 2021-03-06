// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	common "github.com/krishnarajvr/micro-common"
	mock "github.com/stretchr/testify/mock"

	model "micro/module/user/model"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: form
func (_m *IUserRepo) Add(form *model.UserForm) (*model.User, error) {
	ret := _m.Called(form)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.UserForm) *model.User); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserForm) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *IUserRepo) Delete(id int) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *IUserRepo) Get(id int) (*model.User, error) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: page
func (_m *IUserRepo) List(page common.Pagination) (model.Users, error) {
	ret := _m.Called(page)

	var r0 model.Users
	if rf, ok := ret.Get(0).(func(common.Pagination) model.Users); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Pagination) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: form, id
func (_m *IUserRepo) Patch(form *model.UserPatchForm, id int) (*model.User, error) {
	ret := _m.Called(form, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.UserPatchForm, int) *model.User); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserPatchForm, int) error); ok {
		r1 = rf(form, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: form, id
func (_m *IUserRepo) Update(form *model.UserForm, id int) (*model.User, error) {
	ret := _m.Called(form, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.UserForm, int) *model.User); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.UserForm, int) error); ok {
		r1 = rf(form, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
