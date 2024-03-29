// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/krishnarajvr/micro-common"
	mock "github.com/stretchr/testify/mock"

	model "micro/module/client/model"

	token "micro/util/token"
)

// IClientService is an autogenerated mock type for the IClientService type
type IClientService struct {
	mock.Mock
}

// AddCredential provides a mock function with given fields: content, tenantId
func (_m *IClientService) AddCredential(content *model.ClientCredentialForm, tenantId int) (*model.ClientCredentialDto, error) {
	ret := _m.Called(content, tenantId)

	var r0 *model.ClientCredentialDto
	if rf, ok := ret.Get(0).(func(*model.ClientCredentialForm, int) *model.ClientCredentialDto); ok {
		r0 = rf(content, tenantId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClientCredentialDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClientCredentialForm, int) error); ok {
		r1 = rf(content, tenantId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckCredentials provides a mock function with given fields: form
func (_m *IClientService) CheckCredentials(form *model.ClientLoginForm) (*model.ClientCredential, error) {
	ret := _m.Called(form)

	var r0 *model.ClientCredential
	if rf, ok := ret.Get(0).(func(*model.ClientLoginForm) *model.ClientCredential); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClientCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClientLoginForm) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClientCredentialRoles provides a mock function with given fields: clientCredentialId
func (_m *IClientService) GetClientCredentialRoles(clientCredentialId int) ([]string, error) {
	ret := _m.Called(clientCredentialId)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int) []string); ok {
		r0 = rf(clientCredentialId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(clientCredentialId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClientToken provides a mock function with given fields: credential, roles
func (_m *IClientService) GetClientToken(credential *model.ClientCredential, roles []string) (*token.TokenDetails, error) {
	ret := _m.Called(credential, roles)

	var r0 *token.TokenDetails
	if rf, ok := ret.Get(0).(func(*model.ClientCredential, []string) *token.TokenDetails); ok {
		r0 = rf(credential, roles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*token.TokenDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClientCredential, []string) error); ok {
		r1 = rf(credential, roles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCredential provides a mock function with given fields: id
func (_m *IClientService) GetCredential(id int) (*model.ClientCredentialDto, error) {
	ret := _m.Called(id)

	var r0 *model.ClientCredentialDto
	if rf, ok := ret.Get(0).(func(int) *model.ClientCredentialDto); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClientCredentialDto)
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

// ListCredentials provides a mock function with given fields: page
func (_m *IClientService) ListCredentials(page common.Pagination) (model.ClientCredentialDtos, *common.PageResult, error) {
	ret := _m.Called(page)

	var r0 model.ClientCredentialDtos
	if rf, ok := ret.Get(0).(func(common.Pagination) model.ClientCredentialDtos); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.ClientCredentialDtos)
		}
	}

	var r1 *common.PageResult
	if rf, ok := ret.Get(1).(func(common.Pagination) *common.PageResult); ok {
		r1 = rf(page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*common.PageResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(common.Pagination) error); ok {
		r2 = rf(page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PatchCredential provides a mock function with given fields: form, id
func (_m *IClientService) PatchCredential(form *model.ClientCredentialPatchForm, id int) (*model.ClientCredentialDto, error) {
	ret := _m.Called(form, id)

	var r0 *model.ClientCredentialDto
	if rf, ok := ret.Get(0).(func(*model.ClientCredentialPatchForm, int) *model.ClientCredentialDto); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ClientCredentialDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ClientCredentialPatchForm, int) error); ok {
		r1 = rf(form, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshToken provides a mock function with given fields: refreshToken
func (_m *IClientService) RefreshToken(refreshToken string) (*token.TokenDetails, error) {
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
