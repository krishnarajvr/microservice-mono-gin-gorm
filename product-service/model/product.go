package model

import (
	"gorm.io/gorm"
)

type Products []*Product

type Product struct {
	gorm.Model
	Name        string
	Code        string
	Description string
}

func (b Product) ToDto() *ProductDto {
	return &ProductDto{
		ID:          b.ID,
		Name:        b.Name,
		Code:        b.Code,
		Description: b.Description,
	}
}

type ProductDtos []*ProductDto

type ProductDto struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"Product 1"`
	Code        string `json:"code" example:"Product1"`
	Description string `json:"description" example:"Product description"`
}

type ProductForm struct {
	Name        string `json:"name" example:"Product 1" valid:"Required;MaxSize(100)"`
	Code        string `json:"code" example:"Product1" valid:"Required;MaxSize(100)"`
	Description string `json:"description" example:"Product description" valid:"Required;MaxSize(255)"`
}

type ProductPatchForm struct {
	Name        string `json:"name" example:"Product 1" valid:"MaxSize(100)"`
	Code        string `json:"code" example:"Product1" valid:"MaxSize(100)"`
	Description string `json:"description" example:"Product description" valid:"MaxSize(255)"`
}

func (bs Products) ToDto() ProductDtos {
	dtos := make([]*ProductDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

func (f *ProductForm) ToModel() (*Product, error) {
	return &Product{
		Name:        f.Name,
		Code:        f.Code,
		Description: f.Description,
	}, nil
}

func (f *ProductPatchForm) ToModel() (*Product, error) {
	return &Product{
		Name:        f.Name,
		Code:        f.Code,
		Description: f.Description,
	}, nil
}
