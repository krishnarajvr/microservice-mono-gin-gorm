package swagdto

import (
	"micro/dto"
)

type UserResponse struct {
	Status uint         `json:"status" example:"200"`
	Error  interface{}  `json:"error"`
	Data   dto.UserDtos `json:"data"`
}
