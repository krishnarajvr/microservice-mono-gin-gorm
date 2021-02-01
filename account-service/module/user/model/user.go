package model

import (
	"gorm.io/gorm"
)

type Users []*User

type User struct {
	gorm.Model
	Name   string
	Code   string
	Email  string
	Domain string
	Secret string
}

func (b User) ToDto() *UserDto {
	return &UserDto{
		ID:     b.ID,
		Name:   b.Name,
		Code:   b.Code,
		Domain: b.Domain,
		Email:  b.Email,
	}
}

type UserDtos []*UserDto

type UserDto struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"User 1"`
	Domain string `json:"domain" example:"EBOOK"`
	Code   string `json:"code" example:"User1"`
	Email  string `json:"email" example:"user@mail.com"`
	Secret string `json:"secret"`
}

type UserForm struct {
	Name   string `json:"name" example:"User 1" valid:"Required;MaxSize(100)"`
	Code   string `json:"code" example:"User1" valid:"Required;MaxSize(100)"`
	Domain string `json:"domain" example:"EBOOK" valid:"Required;MaxSize(100)"`
	Email  string `json:"email" example:"user@mail.com" valid:"Required;MaxSize(255)"`
	Secret string `json:"secret"`
}

type UserPatchForm struct {
	Name   string `json:"name" example:"User 1" valid:"MaxSize(100)"`
	Code   string `json:"code" example:"User1" valid:"MaxSize(100)"`
	Domain string `json:"code" example:"EBOOK" valid:"MaxSize(100)"`
	Email  string `json:"email" example:"user@mail.com" valid:"MaxSize(255)"`
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
		Name:   f.Name,
		Code:   f.Code,
		Email:  f.Email,
		Secret: "123",
	}, nil
}

func (f *UserPatchForm) ToModel() (*User, error) {
	return &User{
		Name:  f.Name,
		Code:  f.Code,
		Email: f.Email,
	}, nil
}
