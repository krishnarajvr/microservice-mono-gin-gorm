// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	common "github.com/krishnarajvr/micro-common"

	mock "github.com/stretchr/testify/mock"

	model "micro/module/user/model"

	token "micro/util/token"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// Add provides a mock function with given fields: form
func (_m *IUserService) Add(form *model.UserForm) (*model.UserDto, error) {
	ret := _m.Called(form)

	var r0 *model.UserDto
	if rf, ok := ret.Get(0).(func(*model.UserForm) *model.UserDto); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDto)
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

// CheckPermission provides a mock function with given fields: role, module, resource, method
func (_m *IUserService) CheckPermission(role string, module string, resource string, method string) (bool, error) {
	ret := _m.Called(role, module, resource, method)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, string, string) bool); ok {
		r0 = rf(role, module, resource, method)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(role, module, resource, method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *IUserService) Delete(id int) (*model.UserDto, error) {
	ret := _m.Called(id)

	var r0 *model.UserDto
	if rf, ok := ret.Get(0).(func(int) *model.UserDto); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDto)
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

// ExtractTokenMetadata provides a mock function with given fields: r
func (_m *IUserService) ExtractTokenMetadata(r *http.Request) (*token.AccessDetails, error) {
	ret := _m.Called(r)

	var r0 *token.AccessDetails
	if rf, ok := ret.Get(0).(func(*http.Request) *token.AccessDetails); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*token.AccessDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: tenantId, id
func (_m *IUserService) Get(tenantId int, id int) (*model.UserDto, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.UserDto
	if rf, ok := ret.Get(0).(func(int, int) *model.UserDto); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDto)
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

// GetToken provides a mock function with given fields: user
func (_m *IUserService) GetToken(user *model.User) (*token.TokenDetails, error) {
	ret := _m.Called(user)

	var r0 *token.TokenDetails
	if rf, ok := ret.Get(0).(func(*model.User) *token.TokenDetails); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*token.TokenDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: tenantId, page, filters
func (_m *IUserService) List(tenantId int, page common.Pagination, filters model.UserFilterList) (model.UserDtos, *common.PageResult, error) {
	ret := _m.Called(tenantId, page, filters)

	var r0 model.UserDtos
	if rf, ok := ret.Get(0).(func(int, common.Pagination, model.UserFilterList) model.UserDtos); ok {
		r0 = rf(tenantId, page, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.UserDtos)
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

// Login provides a mock function with given fields: login
func (_m *IUserService) Login(login *model.LoginForm) (*model.User, error) {
	ret := _m.Called(login)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.LoginForm) *model.User); ok {
		r0 = rf(login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.LoginForm) error); ok {
		r1 = rf(login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseUri provides a mock function with given fields: Uri
func (_m *IUserService) ParseUri(Uri string) (map[string]string, bool) {
	ret := _m.Called(Uri)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(Uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(Uri)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: form, id
func (_m *IUserService) Patch(form *model.UserPatchForm, id int) (*model.UserDto, error) {
	ret := _m.Called(form, id)

	var r0 *model.UserDto
	if rf, ok := ret.Get(0).(func(*model.UserPatchForm, int) *model.UserDto); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDto)
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

// RefreshToken provides a mock function with given fields: refreshToken
func (_m *IUserService) RefreshToken(refreshToken string) (*token.TokenDetails, error) {
	ret := _m.Called(refreshToken)

	var r0 *token.TokenDetails
	if rf, ok := ret.Get(0).(func(string) *token.TokenDetails); ok {
		r0 = rf(refreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*token.TokenDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: form, id
func (_m *IUserService) Update(form *model.UserForm, id int) (*model.UserDto, error) {
	ret := _m.Called(form, id)

	var r0 *model.UserDto
	if rf, ok := ret.Get(0).(func(*model.UserForm, int) *model.UserDto); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserDto)
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
