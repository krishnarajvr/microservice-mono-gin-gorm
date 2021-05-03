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
	Password  string
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
	FirstName string `json:"firstName" example:"John" valid:"Required;MinSize(2);MaxSize(255)"`
	LastName  string `json:"lastName" example:"Doe" valid:"MaxSize(255)"`
	Password  string `json:"password" example:"Pass!23" valid:"MinSize(5);MaxSize(20)"`
}

type LoginForm struct {
	Email    string `json:"email" label:"email" example:"john@mail.com" valid:"Required;Email;"`
	Password string `json:"password" label:"password"  example:"john123" valid:"MinSize(5);MaxSize(20)"`
}

type UserPatchForm struct {
	Name      string `json:"name" example:"User 1" valid:"MinSize(2);MaxSize(255)"`
	Email     string `json:"email" example:"john@mail.com" valid:"Email;"`
	FirstName string `json:"firstName" example:"John" valid:"MinSize(2);MaxSize(255)"`
	LastName  string `json:"lastName" example:"Doe" valid:"MaxSize(255)"`
}

type UserFilterList struct {
	Name  string `json:"name" form:"name" example:"John"`
	Email string `json:"email" form:"email" example:"john@mail.com"`
}

type AuthorizeData struct {
	Uri    string `json:"uri" form:"uri" valid:"MinSize(2);`
	Method string `json:"method" form:"method" valid:"MinSize(2);MaxSize(10)"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImIyY2Fik5MmItNGZiZi1.v_4EzMc1HG9mJZCGNk0UKnqTveOAjtgIO9Za4tHDBY"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGcIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzNjA2MzQsInJlZfdXVpZCI6ImIyY2FyZWMzNnVzZXJfaWQiOjF9.hv_4EzMc1HG9mJZCGNk0UKnqTveOAjtgIO9Za4tHDBY"`
}

type TokenRefresh struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGcIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzNjA2MzQsInJlZfdXVpZCI6ImIyY2FyZWMzNnVzZXJfaWQiOjF9.hv_4EzMc1HG9mJZCGNk0UKnqTveOAjtgIO9Za4tHDBY"`
}

type AuthRespose struct {
	TenantId    string   `json:"tenantId"`
	UserId      string   `json:"userId"`
	ReferenceId string   `json:"referenceId"`
	Type        string   `json:"type"`
	Subject     string   `json:"subject"`
	Roles       []string `json:"roles"`
	Admin       bool     `json:"admin"`
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
