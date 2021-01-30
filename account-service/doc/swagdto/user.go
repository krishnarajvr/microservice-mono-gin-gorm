package swagdto

import (
	"micro/model"
)

type UserResponse struct {
	Status uint           `json:"status" example:"200"`
	Error  interface{}    `json:"error"`
	Data   model.UserDtos `json:"data"`
}
