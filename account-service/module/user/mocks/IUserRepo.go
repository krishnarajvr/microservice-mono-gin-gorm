// Code generated by mockery v1.0.0. DO NOT EDIT.

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
func (_m *IUserRepo) Add(form *model.User) (*model.User, error) {
	ret := _m.Called(form)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
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

// Get provides a mock function with given fields: tenantId, id
func (_m *IUserRepo) Get(tenantId int, id int) (*model.User, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(int, int) *model.User); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(tenantId, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *IUserRepo) GetByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: tenantId, page, filters
func (_m *IUserRepo) List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.Users, *common.PageResult, error) {
	ret := _m.Called(tenantId, page, filters)

	var r0 model.Users
	if rf, ok := ret.Get(0).(func(int, common.Pagination, model.UserFilterList) model.Users); ok {
		r0 = rf(tenantId, page, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Users)
		}
	}

	var r1 *common.PageResult
	if rf, ok := ret.Get(1).(func(int, common.Pagination, model.UserFilterList) *common.PageResult); ok {
		r1 = rf(tenantId, page, filters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*common.PageResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, common.Pagination, model.UserFilterList) error); ok {
		r2 = rf(tenantId, page, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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