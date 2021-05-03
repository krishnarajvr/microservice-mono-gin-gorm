package model

import (
	"time"

	"gorm.io/gorm"
)

type Tenants []*Tenant

type Tenant struct {
	gorm.Model
	Name      string
	Code      string
	Email     string
	Domain    string
	Secret    string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TenantDtos []*TenantDto

type TenantDto struct {
	ID        uint      `json:"id" example:"123"`
	Name      string    `json:"name" example:"Tenant 1"`
	Domain    string    `json:"domain" example:"EBOOK"`
	Code      string    `json:"code" example:"Tenant1"`
	Email     string    `json:"email" example:"tenant@mail.com"`
	IsActive  bool      `json:"isActive" example:true`
	CreatedAt time.Time `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type TenantForm struct {
	Name   string `json:"name" example:"Tenant 1" valid:"Required;MinSize(5);MaxSize(255)"`
	Code   string `json:"code" example:"Tenant1" valid:"Required;MinSize(3);MaxSize(50)"`
	Domain string `json:"domain" example:"EBOOK" valid:"Required;MinSize(3);MaxSize(50)"`
	Email  string `json:"email" example:"tenant@mail.com" valid:"Required;Email;"`
}

type TenantPatchForm struct {
	Name   string `json:"name" example:"Tenant 1" valid:"MinSize(5);MaxSize(255)"`
	Code   string `json:"code" example:"Tenant1" valid:"MinSize(3);MaxSize(50"`
	Domain string `json:"domain" example:"EBOOK" valid:"Required;MinSize(3);MaxSize(50)"`
	Email  string `json:"email" example:"tenant@mail.com" valid:"Email"`
}

type TenantRegisterForm struct {
	Name      string `json:"name" example:"Tenant1" valid:"Required;MinSize(5);MaxSize(255)"`
	Password  string `json:"password" example:"Pass@1" valid:"Required;MinSize(5);MaxSize(255)"`
	Domain    string `json:"domain" example:"eBook" valid:"Required;MinSize(3);MaxSize(50)"`
	Email     string `json:"email" example:"tenant1@mail.com" valid:"Required;Email;"`
	FirstName string `json:"firstName" example:"John" valid:"MinSize(2);MaxSize(255)"`
	LastName  string `json:"lastName" example:"Doe" valid:"MaxSize(255)"`
}

func (f *TenantForm) ToModel() (*Tenant, error) {
	return &Tenant{
		Name:   f.Name,
		Code:   f.Code,
		Email:  f.Email,
		Domain: f.Domain,
	}, nil
}

func (f *TenantPatchForm) ToModel() (*Tenant, error) {
	return &Tenant{
		Name:  f.Name,
		Code:  f.Code,
		Email: f.Email,
	}, nil
}

func (b Tenant) ToDto() *TenantDto {
	return &TenantDto{
		ID:        b.ID,
		Name:      b.Name,
		Code:      b.Code,
		Domain:    b.Domain,
		Email:     b.Email,
		IsActive:  b.IsActive,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func (bs Tenants) ToDto() TenantDtos {
	dtos := make([]*TenantDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

func (b TenantRegisterForm) ToTenantForm() *TenantForm {
	return &TenantForm{
		Name:   b.Name,
		Domain: b.Domain,
		Email:  b.Email,
	}
}
