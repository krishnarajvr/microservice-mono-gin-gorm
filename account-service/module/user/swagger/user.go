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
