package swagdto

import (
	"micro/model"
)

type ProductSampleData struct {
	ProductData model.ProductDto `json:"product"`
}

type ProductSampleListData struct {
	ProductData model.ProductDtos `json:"products"`
}

type ProductListResponse struct {
	Status uint                  `json:"status" example:"200"`
	Error  interface{}           `json:"error"`
	Data   ProductSampleListData `json:"data"`
}

type ProductResponse struct {
	Status uint              `json:"status" example:"200"`
	Error  interface{}       `json:"error"`
	Data   ProductSampleData `json:"data"`
}
