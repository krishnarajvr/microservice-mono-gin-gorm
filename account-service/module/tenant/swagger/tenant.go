package swagger

import (
	"micro/module/tenant/model"
)

type TenantSampleData struct {
	TenantData model.TenantDto `json:"tenant"`
}

type TenantSampleListData struct {
	TenantData model.TenantDtos `json:"tenants"`
}

type TenantListResponse struct {
	Status uint                 `json:"status" example:"200"`
	Error  interface{}          `json:"error"`
	Data   TenantSampleListData `json:"data"`
}

type TenantResponse struct {
	Status uint             `json:"status" example:"200"`
	Error  interface{}      `json:"error"`
	Data   TenantSampleData `json:"data"`
}
