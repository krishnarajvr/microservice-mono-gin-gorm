package model

import (
	"time"

	"gorm.io/gorm"
)

type Users []*User

type User struct {
	gorm.Model
	Name        string
	Email       string
	CreatedDate time.Time
	FirstName   string
}

func (b User) ToDto() *UserDto {
	return &UserDto{
		ID:    b.ID,
		Name:  b.Name,
		Email: b.Email,
	}
}

type UserDtos []*UserDto

type UserDto struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"bob"`
	Email       string `json:"email" example:"bob@gmail.com"`
	CreatedDate string `json:"created_date"`
	FirstName   string `json:"firstName" example:"Bob"`
}

type UserForm struct {
	Name        string `json:"Name" form:"required,max=255"`
	Email       string `json:"Email" form:"required,alpha_space,max=255"`
	CreatedDate string `json:"published_date" form:"required,date"`
	FirstName   string `json:"image_url" form:"url"`
}

func (bs Users) ToDto() UserDtos {
	dtos := make([]*UserDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

func (f *UserForm) ToModel() (*User, error) {
	pubDate, err := time.Parse("2006-01-02", f.CreatedDate)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:        f.Name,
		Email:       f.Email,
		CreatedDate: pubDate,
		FirstName:   f.FirstName,
	}, nil
}
