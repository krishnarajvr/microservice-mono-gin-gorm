package model

import (
	"gorm.io/gorm"
)

type Tenants []*Tenant

type Tenant struct {
	gorm.Model
	Name   string
	Code   string
	Email  string
	Domain string
	Secret string
}

func (b Tenant) ToDto() *TenantDto {
	return &TenantDto{
		ID:     b.ID,
		Name:   b.Name,
		Code:   b.Code,
		Domain: b.Domain,
		Email:  b.Email,
	}
}

type TenantDtos []*TenantDto

type TenantDto struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"Tenant 1"`
	Domain string `json:"domain" example:"EBOOK"`
	Code   string `json:"code" example:"Tenant1"`
	Email  string `json:"email" example:"tenant@mail.com"`
	Secret string `json:"secret"`
}

type TenantForm struct {
	Name   string `json:"name" example:"Tenant 1" valid:"Required;MaxSize(100)"`
	Code   string `json:"code" example:"Tenant1" valid:"Required;MaxSize(100)"`
	Domain string `json:"domain" example:"EBOOK" valid:"Required;MaxSize(100)"`
	Email  string `json:"email" example:"tenant@mail.com" valid:"Required;MaxSize(255)"`
	Secret string `json:"secret"`
}

type TenantPatchForm struct {
	Name   string `json:"name" example:"Tenant 1" valid:"MaxSize(100)"`
	Code   string `json:"code" example:"Tenant1" valid:"MaxSize(100)"`
	Domain string `json:"code" example:"EBOOK" valid:"MaxSize(100)"`
	Email  string `json:"email" example:"tenant@mail.com" valid:"MaxSize(255)"`
}

func (bs Tenants) ToDto() TenantDtos {
	dtos := make([]*TenantDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

func (f *TenantForm) ToModel() (*Tenant, error) {
	return &Tenant{
		Name:   f.Name,
		Code:   f.Code,
		Email:  f.Email,
		Secret: "123",
	}, nil
}

func (f *TenantPatchForm) ToModel() (*Tenant, error) {
	return &Tenant{
		Name:  f.Name,
		Code:  f.Code,
		Email: f.Email,
	}, nil
}
