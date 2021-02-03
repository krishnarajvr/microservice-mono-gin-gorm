package model

import (
	"time"

	"gorm.io/gorm"
)

type Users []*User

type User struct {
	gorm.Model
	TenantId  uint
	Name      string
	Email     string
	FirstName string
	LastName  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDtos []*UserDto

type UserDto struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"User 1"`
	Email     string    `json:"email" example:"user@mail.com"`
	FirstName string    `json:"firstName" example:"John"`
	LastName  string    `json:"lastName" example:"Doe"`
	IsActive  bool      `json:"isActive" example:true`
	CreatedAt time.Time `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

func (b User) ToDto() *UserDto {
	return &UserDto{
		ID:        b.ID,
		Name:      b.Name,
		Email:     b.Email,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		IsActive:  b.IsActive,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

type UserForm struct {
	Name      string `json:"name" example:"User 1" valid:"Required;MinSize(2);MaxSize(255)"`
	Email     string `json:"email" example:"john@mail.com" valid:"Required;Email;"`
	FirstName string `json:"code" example:"John" valid:"Required;MinSize(2);MaxSize(255)"`
	LastName  string `json:"domain" example:"Doe" valid:"MaxSize(255)"`
}

type UserPatchForm struct {
	Name      string `json:"name" example:"User 1" valid:"MinSize(2);MaxSize(255)"`
	Email     string `json:"email" example:"john@mail.com" valid:"Email;"`
	FirstName string `json:"code" example:"John" valid:"MinSize(2);MaxSize(255)"`
	LastName  string `json:"domain" example:"Doe" valid:"MaxSize(255)"`
}

func (bs Users) ToDto() UserDtos {
	dtos := make([]*UserDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

func (f *UserForm) ToModel() (*User, error) {
	return &User{
		Name:      f.Name,
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
	}, nil
}

func (f *UserPatchForm) ToModel() (*User, error) {
	return &User{
		Name:      f.Name,
		Email:     f.Email,
		FirstName: f.FirstName,
		LastName:  f.LastName,
	}, nil
}
