package dto

import (
	"time"
)

type UserDtos []*UserDto

type UserDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_date"`
	FirstName   string `json:"firstName"`
}

func (b User) ToDto() *UserDto {
	return &UserDto{
		ID:          b.ID,
		Name:        b.Name,
		Email:       b.Email,
		CreatedDate: b.CreatedDate.Format("2006-01-02"),
		FirstName:   b.FirstName,
	}
}

func (bs Users) ToDto() UserDtos {
	dtos := make([]*UserDto, len(bs))
	for i, b := range bs {
		dtos[i] = b.ToDto()
	}

	return dtos
}

type UserForm struct {
	Name        string `json:"Name" form:"required,max=255"`
	Email       string `json:"Email" form:"required,alpha_space,max=255"`
	CreatedDate string `json:"published_date" form:"required,date"`
	FirstName   string `json:"image_url" form:"url"`
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
