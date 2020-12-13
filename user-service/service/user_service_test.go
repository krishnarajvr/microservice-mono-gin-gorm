package service

import (
	"testing"

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

		IUserRepo := new(mocks.IUserRepo)

		IUserRepo.On("ListUsers").Return(users, nil)

		us := NewUserService(&ServiceConfig{
			UserRepo: IUserRepo,
		})

		u, err := us.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, u, userDtos)
		IUserRepo.AssertExpectations(t)
	})

}
