package swagger

import (
	"time"
)

type ClientCredentialDtos []*ClientCredentialDto

type ClientCredentialDto struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"client 1"`
	Code        string    `json:"code" example:"client 1"`
	Secret      string    `json:"secret" example:"12-345-567"`
	ReferenceId string    `json:"referenceId" example:"user id"`
	Type        string    `json:"type" example:"user"`
	Payload     Payload   `json:"payload" `
	IsActive    int       `json:"isActive" example:"1"`
	CreatedAt   time.Time `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type ClientTokenResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NXVCJ9.eyJhY2Nlc30IDIiLCJ0eXBlIjoiY2xpZW50In0.XBjAxzruIT"`
	TokenType    string `json:"token_type" example:"bearer" `
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI54XVCJ54.eyJhY2Nlc30IDIiLCJ0eXBlIjoiY2xpZW5045.XBjAxzru45"`
}

type ClientLoginForm struct {
	Type   string `json:"grant_type" label:"grant_type" form:"grant_type" example:"client_credentials" valid:"Required;MinSize(2);MaxSize(255)"`
	Code   string `json:"client_id" label:"client_id" form:"client_id" example:"ClientA" valid:"Required;MinSize(1);MaxSize(255)"`
	Secret string `json:"client_secret"  label:"client_secret" form:"client_secret" example:"@#WERWREWRWERWE" valid:"Required;MinSize(1);MaxSize(255)"`
}

type ClientCredentialSampleData struct {
	ClientCredentialData ClientCredentialDto `json:"client"`
}

type ClientCredentialSampleListData struct {
	ClientCredentialData ClientCredentialDtos `json:"client"`
}

type ClientListCredentialsResponse struct {
	Status uint                           `json:"status" example:"200"`
	Error  interface{}                    `json:"error"`
	Data   ClientCredentialSampleListData `json:"data"`
}

type ClientCredentialResponse struct {
	Status uint                       `json:"status" example:"200"`
	Error  interface{}                `json:"error"`
	Data   ClientCredentialSampleData `json:"data"`
}

type ClientCredentialForm struct {
	Name        string  `json:"name" example:"Client A" `
	Code        string  `json:"code" example:"ClientA" `
	ReferenceId string  `json:"referenceId" example:"client_id" `
	Type        string  `json:"type" example:"user"`
	Payload     Payload `json:"payload" `
	IsActive    int     `json:"isActive" example:"1" `
}

type Payload struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}
