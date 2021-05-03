package service

import (
	"testing"

	"github.com/krishnarajvr/micro-common/locale"
	"micro/config"
	"micro/module/tenant/mocks"
	"micro/module/tenant/model"

	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
)

func TestTenant(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockTenantResp := &model.Tenant{
			Name: "Tenant 1",
			Code: "Code 1",
		}

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		products := model.Tenants{mockTenantResp}
		productDtos := products.ToDto()

		ITenantRepo := new(mocks.ITenantRepo)

		ITenantRepo.On("List", page).Return(products, nil)

		appConf := config.AppConfig()
		langLocale := locale.Locale{}
		lang := langLocale.New(appConf.App.Lang)

		ps := NewService(ServiceConfig{
			TenantRepo: ITenantRepo,
			Lang:       lang,
		})

		u, err := ps.List(page)

		assert.NoError(t, err)
		assert.Equal(t, u, productDtos)
		ITenantRepo.AssertExpectations(t)
	})

}
