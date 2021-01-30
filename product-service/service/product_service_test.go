package service

import (
	"testing"

	"micro/app/locale"
	"micro/config"
	"micro/model"
	"micro/repo/mocks"

	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockProductResp := &model.Product{
			Name:        "Product 1",
			Code:        "Code 1",
			Description: "Description 1",
		}

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		products := model.Products{mockProductResp}
		productDtos := products.ToDto()

		IProductRepo := new(mocks.IProductRepo)

		IProductRepo.On("List", page).Return(products, nil)

		appConf := config.AppConfig()
		langLocale := locale.Locale{}
		lang := langLocale.New(appConf.App.Lang)

		ps := NewProductService(&ServiceConfig{
			ProductRepo: IProductRepo,
			Lang:        lang,
		})

		u, err := ps.List(page)

		assert.NoError(t, err)
		assert.Equal(t, u, productDtos)
		IProductRepo.AssertExpectations(t)
	})

}
