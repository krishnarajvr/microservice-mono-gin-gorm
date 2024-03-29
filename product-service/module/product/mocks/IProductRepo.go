// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/krishnarajvr/micro-common"
	mock "github.com/stretchr/testify/mock"

	model "micro/module/product/model"
)

// IProductRepo is an autogenerated mock type for the IProductRepo type
type IProductRepo struct {
	mock.Mock
}

// Add provides a mock function with given fields: form
func (_m *IProductRepo) Add(form *model.Product) (*model.Product, error) {
	ret := _m.Called(form)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(*model.Product) *model.Product); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Product) error); ok {
		r1 = rf(form)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: tenantId, id
func (_m *IProductRepo) Delete(tenantId int, id int) (*model.Product, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(int, int) *model.Product); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
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

// Get provides a mock function with given fields: tenantId, id
func (_m *IProductRepo) Get(tenantId int, id int) (*model.ProductDetail, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.ProductDetail
	if rf, ok := ret.Get(0).(func(int, int) *model.ProductDetail); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductDetail)
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

// List provides a mock function with given fields: tenantId, page, filters
func (_m *IProductRepo) List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductsList, *common.PageResult, error) {
	ret := _m.Called(tenantId, page, filters)

	var r0 model.ProductsList
	if rf, ok := ret.Get(0).(func(int, common.Pagination, model.ProductFilterList) model.ProductsList); ok {
		r0 = rf(tenantId, page, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.ProductsList)
		}
	}

	var r1 *common.PageResult
	if rf, ok := ret.Get(1).(func(int, common.Pagination, model.ProductFilterList) *common.PageResult); ok {
		r1 = rf(tenantId, page, filters)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*common.PageResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, common.Pagination, model.ProductFilterList) error); ok {
		r2 = rf(tenantId, page, filters)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Patch provides a mock function with given fields: form, id
func (_m *IProductRepo) Patch(form *model.ProductPatchForm, id int) (*model.Product, error) {
	ret := _m.Called(form, id)

	var r0 *model.Product
	if rf, ok := ret.Get(0).(func(*model.ProductPatchForm, int) *model.Product); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.ProductPatchForm, int) error); ok {
		r1 = rf(form, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
