package swagger

import (
	"micro/module/user/model"
)

type UserSampleData struct {
	UserData model.UserDto `json:"user"`
}

type UserSampleListData struct {
	UserData model.UserDtos `json:"users"`
}

type UserListResponse struct {
	Status uint               `json:"status" example:"200"`
	Error  interface{}        `json:"error"`
	Data   UserSampleListData `json:"data"`
}

type UserResponse struct {
	Status uint           `json:"status" example:"200"`
	Error  interface{}    `json:"error"`
	Data   UserSampleData `json:"data"`
}

type UserTokenModel struct {
	TokenType    string `json:"tokenType" example:"bearer" `
	AccessToken  string `json:"accessToken" example:"eyJhbGciOiJIUzI1NXVCJ9.eyJhY2Nlc30IDIiLCJ0eXBlIjoiY2xpZW50In0.XBjAxzruIT"`
	RefreshToken string `json:"refreshToken" example:"eyJhbGciOiJIUzI54XVCJ54.eyJhY2Nlc30IDIiLCJ0eXBlIjoiY2xpZW5045.XBjAxzru45"`
	AtExpires    string `json:"accessTokenExpiry" example:"2021-04-24T08:40:59Z"`
	RtExpires    string `json:"refreshTokenExpiry" example:"2021-04-21T12:40:59Z"`
}

type UserTokenData struct {
	TokenData UserTokenModel `json:"token"`
}

type TokenResponse struct {
	Status    uint          `json:"status" example:"200"`
	Error     interface{}   `json:"error"`
	Data      UserTokenData `json:"data"`
	RequestId string        `json:"requestId" example:"3b6272b9-1ef1-45e0"`
}
