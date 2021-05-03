package swagger

import (
	"time"
)

type ProductDtos []*ProductDto

type ProductDto struct {
	ID          uint        `json:"id" example:"1"`
	Name        string      `json:"name" example:"Product 1"`
	Code        string      `json:"code" example:"10011" `
	Description string      `json:"description" example:"product description"`
	Meta        ProductMeta `json:"meta"`
	CreatedAt   time.Time   `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time   `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type ProductListDtos []*ProductListDto

type ProductListDto struct {
	ID          uint        `json:"id" example:"1"`
	Name        string      `json:"name" example:"Product 1"`
	Code        string      `json:"code" example:"10011" `
	Description string      `json:"description" example:"product description"`
	Meta        ProductMeta `json:"meta"`
	CreatedAt   time.Time   `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time   `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type ProductDetailDto struct {
	ID          uint        `json:"id" example:"1"`
	Name        string      `json:"name" example:"Product 1"`
	Code        string      `json:"code" example:"10011" `
	Description string      `json:"description" example:"product description"`
	Meta        ProductMeta `json:"meta"`
	CreatedAt   time.Time   `json:"createdAt" example:"2021-02-02T02:52:24Z"`
	UpdatedAt   time.Time   `json:"updatedAt" example:"2021-02-02T02:52:24Z"`
}

type ProductAddData struct {
	ProductData ProductDto `json:"product"`
}

type ProductupdateData struct {
	ProductData ProductDto `json:"product"`
}

type ProductSampleData struct {
	ProductData ProductDetailDto `json:"product"`
}

type ProductSampleListData struct {
	ProductData ProductListDtos `json:"product"`
}
type ProductAddResponse struct {
	Status uint           `json:"status" example:"200"`
	Error  interface{}    `json:"error"`
	Data   ProductAddData `json:"data"`
}

type ProductUpdateResponse struct {
	Status uint              `json:"status" example:"200"`
	Error  interface{}       `json:"error"`
	Data   ProductupdateData `json:"data"`
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

type ProductForm struct {
	Name        string      `json:"name" example:"Product 1" valid:"Required;MinSize(2);MaxSize(255)"`
	Code        string      `json:"code" example:"1001"  valid:"Required;AlphaNumeric;MinSize(1);MaxSize(255)"`
	Description string      `json:"description" example:"product description"`
	Meta        ProductMeta `json:"meta"`
	IsActive    int         `json:"isActive" example:"1"`
}

type ProductPatchForm struct {
	Name        string      `json:"name" example:"Product 1" valid:"Required;MinSize(2);MaxSize(255)"`
	Code        string      `json:"code" example:"1001"  valid:"Required;AlphaNumeric;MinSize(1);MaxSize(255)"`
	Description string      `json:"description" example:"product description"`
	Meta        ProductMeta `json:"meta"`
}

type ProductMeta struct {
	Author  string `json:"author" example:"AuthorA"`
	Country string `json:"country" example:"USA"`
}
