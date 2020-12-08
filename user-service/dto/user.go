package dto

import (
	"micro/model"
	"time"
)

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

func (f *UserForm) ToModel() (*model.User, error) {
	pubDate, err := time.Parse("2006-01-02", f.CreatedDate)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Name:        f.Name,
		Email:       f.Email,
		CreatedDate: pubDate,
		FirstName:   f.FirstName,
	}, nil
}
