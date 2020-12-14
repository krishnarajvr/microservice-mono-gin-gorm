package service

import (
	"testing"

	"micro/model"
	"micro/repo/mocks"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockProductResp := &model.Product{
			Name:      "krishna",
			FirstName: "krishnaraj.vr@gmail.com",
		}
		products := model.Products{mockProductResp}
		productDtos := products.ToDto()

		IProductRepo := new(mocks.IProductRepo)

		IProductRepo.On("ListProducts").Return(products, nil)

		us := NewProductService(&ServiceConfig{
			ProductRepo: IProductRepo,
		})

		u, err := us.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, u, productDtos)
		IProductRepo.AssertExpectations(t)
	})

}
