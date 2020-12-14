package model

import (
	"time"

	"gorm.io/gorm"
)

type Products []*Product

type Product struct {
	gorm.Model
	Name        string
	Code        string
	CreatedDate time.Time
	Description string
}

func (b Product) ToDto() *ProductDto {
	return &ProductDto{
		ID:   b.ID,
		Name: b.Name,
		Code: b.Code,
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
	Name        string `json:"name" example:"Product 1"`
	Code        string `json:"code" example:"Product1"`
	Description string `json:"description" example:"Product description"`
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
