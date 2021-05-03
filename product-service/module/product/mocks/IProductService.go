// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/krishnarajvr/micro-common"
	datatypes "gorm.io/datatypes"

	mock "github.com/stretchr/testify/mock"

	model "micro/module/product/model"
)

// IProductService is an autogenerated mock type for the IProductService type
type IProductService struct {
	mock.Mock
}

// Add provides a mock function with given fields: tenantId, content
func (_m *IProductService) Add(tenantId int, content *model.ProductForm) (*model.ProductDto, error) {
	ret := _m.Called(tenantId, content)

	var r0 *model.ProductDto
	if rf, ok := ret.Get(0).(func(int, *model.ProductForm) *model.ProductDto); ok {
		r0 = rf(tenantId, content)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *model.ProductForm) error); ok {
		r1 = rf(tenantId, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckProductSchema provides a mock function with given fields: tenantId, jsonData, validateFor
func (_m *IProductService) CheckProductSchema(tenantId int, jsonData []byte, validateFor string) (map[string]interface{}, *map[string]interface{}, error) {
	ret := _m.Called(tenantId, jsonData, validateFor)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(int, []byte, string) map[string]interface{}); ok {
		r0 = rf(tenantId, jsonData, validateFor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 *map[string]interface{}
	if rf, ok := ret.Get(1).(func(int, []byte, string) *map[string]interface{}); ok {
		r1 = rf(tenantId, jsonData, validateFor)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, []byte, string) error); ok {
		r2 = rf(tenantId, jsonData, validateFor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ConvertMapToJson provides a mock function with given fields: metaMap
func (_m *IProductService) ConvertMapToJson(metaMap *map[string]interface{}) (datatypes.JSON, error) {
	ret := _m.Called(metaMap)

	var r0 datatypes.JSON
	if rf, ok := ret.Get(0).(func(*map[string]interface{}) datatypes.JSON); ok {
		r0 = rf(metaMap)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datatypes.JSON)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*map[string]interface{}) error); ok {
		r1 = rf(metaMap)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: tenantId, id
func (_m *IProductService) Delete(tenantId int, id int) (*model.ProductDto, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.ProductDto
	if rf, ok := ret.Get(0).(func(int, int) *model.ProductDto); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductDto)
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
func (_m *IProductService) Get(tenantId int, id int) (*model.ProductDetailDto, error) {
	ret := _m.Called(tenantId, id)

	var r0 *model.ProductDetailDto
	if rf, ok := ret.Get(0).(func(int, int) *model.ProductDetailDto); ok {
		r0 = rf(tenantId, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductDetailDto)
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
func (_m *IProductService) List(tenantId int, page common.Pagination, filters model.ProductFilterList) (model.ProductListDtos, *common.PageResult, error) {
	ret := _m.Called(tenantId, page, filters)

	var r0 model.ProductListDtos
	if rf, ok := ret.Get(0).(func(int, common.Pagination, model.ProductFilterList) model.ProductListDtos); ok {
		r0 = rf(tenantId, page, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.ProductListDtos)
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
func (_m *IProductService) Patch(form *model.ProductPatchForm, id int) (*model.ProductDto, error) {
	ret := _m.Called(form, id)

	var r0 *model.ProductDto
	if rf, ok := ret.Get(0).(func(*model.ProductPatchForm, int) *model.ProductDto); ok {
		r0 = rf(form, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ProductDto)
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