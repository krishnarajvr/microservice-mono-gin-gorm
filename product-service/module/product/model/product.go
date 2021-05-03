package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Products []*Product

type Product struct {
	gorm.Model
	TenantId    int
	Name        string
	Code        string
	Description string
	Meta        datatypes.JSON
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductForm struct {
	Name        string         `json:"name" example:"John" label:"Name" valid:"Required;MinSize(2);MaxSize(255)"`
	Code        string         `json:"code" example:"1001" label:"Code" valid:"Required;AlphaNumeric;MinSize(1);MaxSize(255)"`
	Description string         `json:"description" label:"Description" example:"product description"`
	Meta        datatypes.JSON `json:"meta" label:"Meta" example:"{"businessaccountlead": "smith@gmail.com","technicalcontactlead": "smith@gmail.com"}"`
	IsActive    int            `json:"isActive" valid:"Binary" label:"IsActive" example:"1"`
}

type ProductDtos []*ProductDto

type ProductDto struct {
	ID          uint           `json:"id" example:"1"`
	Name        string         `json:"name" example:"John"`
	Code        string         `json:"code" example:"1001"`
	Description string         `json:"description" example:"product description"`
	Meta        datatypes.JSON `json:"meta" example:"{country:smith@gmail.com,author:smith@gmail.com}"`
	CreatedAt   time.Time      `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time      `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
	IsActive    int            `json:"isActive" valid:"Binary" label:"IsActive" example:"1"`
}

func (v Product) ToDto() *ProductDto {
	return &ProductDto{
		ID:          v.ID,
		Name:        v.Name,
		Code:        v.Code,
		Description: v.Description,
		Meta:        v.Meta,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		IsActive:    v.IsActive,
	}
}

func (vs Products) ToDto() ProductDtos {
	dtos := make([]*ProductDto, len(vs))
	for i, b := range vs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

type ProductPatchForm struct {
	Name        string         `json:"name" example:"John" label:"Name"`
	Code        string         `json:"code" example:"1001" label:"Code"`
	Description string         `json:"description" label:"Description" example:"product description"`
	Meta        datatypes.JSON `json:"meta" label:"Meta" example:"{"country": "smith@gmail.com","author": "smith@gmail.com"}"`
	IsActive    int            `json:"isActive" example:"1" label:"IsActive" valid:"Binary"`
}

func (f *ProductPatchForm) ToModel() (*Product, error) {
	return &Product{
		Name:        f.Name,
		Code:        f.Code,
		Description: f.Description,
		Meta:        f.Meta,
		IsActive:    f.IsActive,
	}, nil
}

func (f *ProductForm) ToModel() (*Product, error) {
	return &Product{
		Name:        f.Name,
		Code:        f.Code,
		Description: f.Description,
		Meta:        f.Meta,
		IsActive:    f.IsActive,
	}, nil
}

type ProductFilterList struct {
	Code string `json:"code" example:"1000"`
	Name string `json:"name" example:"Product 1"`
}

type ProductsList []*ProductList
type ProductListDtos []*ProductListDto

type ProductList struct {
	ID          uint
	Name        string
	Email       string
	Code        string
	Description string
	Meta        datatypes.JSON
	Phone       string
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductListDto struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"John"`
	Code        string    `json:"code" example:"1001"`
	Description string    `json:"description" label:"Description" example:"product description"`
	IsActive    int       `json:"isActive"  example:"1"`
	CreatedAt   time.Time `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type ProductDetail struct {
	ID          uint
	Name        string
	Code        string
	Description string
	Meta        datatypes.JSON
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductDetailDto struct {
	ID          uint           `json:"id" example:"1"`
	Name        string         `json:"name" example:"John"`
	Code        string         `json:"code" example:"1001"`
	Description string         `json:"description" label:"Description" example:"product description"`
	Meta        datatypes.JSON `json:"meta" label:"Meta" example:"{"country": "smith@gmail.com","author": "smith@gmail.com"}"`
	IsActive    int            `json:"isActive"  example:"1"`
	CreatedAt   time.Time      `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time      `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

func (v ProductList) ToProductListDto() *ProductListDto {
	return &ProductListDto{
		ID:          v.ID,
		Name:        v.Name,
		Code:        v.Code,
		Description: v.Description,
		IsActive:    v.IsActive,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}

func (v ProductDetail) ToProductDetailDto() *ProductDetailDto {
	return &ProductDetailDto{
		ID:          v.ID,
		Name:        v.Name,
		Code:        v.Code,
		Description: v.Description,
		IsActive:    v.IsActive,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Meta:        v.Meta,
	}
}

func (bsl ProductsList) ToDto() ProductListDtos {
	dtos := make([]*ProductListDto, len(bsl))
	for i, b := range bsl {
		dtos[i] = b.ToProductListDto()
	}

	return dtos
}
