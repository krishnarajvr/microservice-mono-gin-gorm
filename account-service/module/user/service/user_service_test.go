package service

import (
	"testing"

	"micro/app/locale"
	"micro/config"
	"micro/module/user/mocks"
	"micro/module/user/model"

	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockUserResp := &model.User{
			Name: "User 1",
			Code: "Code 1",
		}

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		products := model.Users{mockUserResp}
		productDtos := products.ToDto()

		IUserRepo := new(mocks.IUserRepo)

		IUserRepo.On("List", page).Return(products, nil)

		appConf := config.AppConfig()
		langLocale := locale.Locale{}
		lang := langLocale.New(appConf.App.Lang)

		ps := NewService(ServiceConfig{
			UserRepo: IUserRepo,
			Lang:     lang,
		})

		u, err := ps.List(page)

		assert.NoError(t, err)
		assert.Equal(t, u, productDtos)
		IUserRepo.AssertExpectations(t)
	})

}
