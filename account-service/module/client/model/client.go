package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ClientCredentials []*ClientCredential

type ClientCredential struct {
	gorm.Model
	TenantId    int
	Name        string
	Code        string
	Secret      string
	ReferenceId string
	Type        string
	Payload     datatypes.JSON
	IsActive    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ClientCredentialRole struct {
	ClientCredentialID uint
	RoleID             uint
}

type ClientCredentialForm struct {
	Name        string         `json:"name" example:"Client A" valid:"Required;MinSize(2);MaxSize(255)"`
	Code        string         `json:"code" example:"ClientA"  valid:"Required;AlphaNumeric;MinSize(1);MaxSize(255)"`
	ReferenceId string         `json:"referenceId" example:"user_id" label:"ReferenceId" valid:"Required;"`
	Type        string         `json:"type" example:"user" label:"Type" valid:"Required"`
	Payload     datatypes.JSON `json:"payload" label:"Payload"`
	IsActive    int            `json:"isActive" example:"1" label:"isActive" valid:"Binary"`
}

type ClientCredentialDtos []*ClientCredentialDto

type ClientCredentialDto struct {
	ID          uint           `json:"id" example:"1"`
	Name        string         `json:"name" example:"client 1"`
	Code        string         `json:"code" example:"1002"`
	Secret      string         `json:"secret" example:"12-345-567"`
	ReferenceId string         `json:"referenceId" example:"user_id"`
	Type        string         `json:"type" example:"user"`
	Payload     datatypes.JSON `json:"payload" example:"{}"`
	IsActive    int            `json:"isActive" example:"1"`
	CreatedAt   time.Time      `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time      `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

func (c ClientCredential) ToCredentialDto() *ClientCredentialDto {
	return &ClientCredentialDto{
		ID:          c.ID,
		Name:        c.Name,
		Code:        c.Code,
		Secret:      c.Secret,
		ReferenceId: c.ReferenceId,
		Type:        c.Type,
		Payload:     c.Payload,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (cc ClientCredentials) ToCredentialDtos() ClientCredentialDtos {
	dtos := make([]*ClientCredentialDto, len(cc))
	for i, b := range cc {
		dtos[i] = b.ToCredentialDto()
	}

	return dtos
}

type ClientCredentialPatchForm struct {
	Name   string `json:"name" example:"Client A" valid:"Required;MinSize(2);MaxSize(255)"`
	Code   string `json:"code" example:"ClientA" valid:"Required;AlphaNumeric;MinSize(1);MaxSize(255)"`
	Secret string `json:"secret" `
	Type   string `json:"type"`
}

type ClientLoginForm struct {
	GrantType    string `json:"grant_type" label:"grant_type" form:"grant_type" example:"client_credentials" valid:"Required;MinSize(2);MaxSize(255)"`
	ClientId     string `json:"client_id" label:"client_id" form:"client_id" example:"ClientA" valid:"Required;MinSize(1);MaxSize(255)"`
	ClientSecret string `json:"client_secret"  label:"client_secret" form:"client_secret" example:"@#WERWREWRWERWE" valid:"Required;MinSize(1);MaxSize(255)"`
	RefreshToken string `json:"refresh_token,omitempty"  label:"refresh_token" form:"refresh_token" example:"@#WERWREWRWERWE" `
}

func (f *ClientCredentialPatchForm) ToModel() (*ClientCredential, error) {
	return &ClientCredential{
		Name:   f.Name,
		Code:   f.Code,
		Secret: f.Secret,
		Type:   f.Type,
	}, nil
}

func (f *ClientCredentialForm) ToModel() (*ClientCredential, error) {
	return &ClientCredential{
		Name:        f.Name,
		Code:        f.Code,
		ReferenceId: f.ReferenceId,
		Type:        f.Type,
		Payload:     f.Payload,
		IsActive:    f.IsActive,
	}, nil
}
