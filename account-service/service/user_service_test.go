package service

import (
	"testing"

	"micro/app/locale"
	"micro/config"
	"micro/model"
	"micro/repo/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockUserResp := &model.User{
			Name:      "krishna",
			FirstName: "krishnaraj.vr@gmail.com",
		}
		users := model.Users{mockUserResp}
		userDtos := users.ToDto()

		userRepo := new(mocks.IUserRepo)

		userRepo.On("ListUsers").Return(users, nil)

		appConf := config.AppConfig()
		langLocale := locale.Locale{}
		lang := langLocale.New(appConf.App.Lang)

		us := NewUserService(&ServiceConfig{
			UserRepo: userRepo,
			Lang:     lang,
		})

		u, err := us.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, u, userDtos)
		userRepo.AssertExpectations(t)
	})

}
