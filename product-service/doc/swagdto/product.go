package swagdto

import (
	"micro/model"
)

type ProductResponse struct {
	Status uint           `json:"status" example:"200"`
	Error  interface{}    `json:"error"`
	Data   model.ProductDtos `json:"data"`
}
