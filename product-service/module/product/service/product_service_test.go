package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"micro/app"
	"micro/config"
	"micro/module/product/mocks"
	"micro/module/product/model"

	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/datatypes"
)

func TestProduct(t *testing.T) {
	t.Run("List Product Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockProductResp := &model.ProductList{
				ID:   1,
				Name: "product 1",
				Code: "code",
			}
			page := common.Pagination{
				PageNumber: "1",
				PageOffset: "1",
				PageSize:   "25",
				Search:     "",
			}

			var mockFilters model.ProductFilterList
			mockFilters.Code = "J10"
			mockFilters.Name = "Sample"

			products := model.ProductsList{mockProductResp}
			productsDtos := products.ToDto()

			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("List", 1, page, mockFilters).Return(products, nil, nil)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, _, err := ps.List(1, page, mockFilters)

			assert.NoError(t, err)
			assert.Equal(t, u, productsDtos)
		})
		t.Run("Error", func(t *testing.T) {
			mockProductResp := &model.ProductList{
				ID:   1,
				Name: "product 1",
				Code: "code",
			}
			page := common.Pagination{
				PageNumber: "1",
				PageOffset: "1",
				PageSize:   "25",
				Search:     "",
			}

			var mockFilters model.ProductFilterList
			mockFilters.Code = "J10"
			mockFilters.Name = "Sample"

			products := model.ProductsList{mockProductResp}
			productsDtos := products.ToDto()

			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("List", 1, page, mockFilters).Return(nil, nil, fmt.Errorf("Some error down call chain"))

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, _, err := ps.List(1, page, mockFilters)
			assert.Error(t, err)
			assert.NotEqual(t, u, productsDtos)
		})
	})

	t.Run("Patch Product Service", func(t *testing.T) {
		t.Run("Patchproduct_200_Success", func(t *testing.T) {
			meta := `{"Author":"Author 1","Country":"USA"}`
			mockProductResp := &model.Product{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participant",
				Meta:        datatypes.JSON(meta),
			}
			mockProductPatchReq := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participant",
				Meta:        datatypes.JSON(meta),
			}

			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("Patch", mock.Anything, 1).Return(mockProductResp, nil, nil)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			_, err := ps.Patch(mockProductPatchReq, 1)

			assert.NoError(t, err)
		})
		t.Run("Patchproduct_500_Failure", func(t *testing.T) {
			meta := map[string]interface{}{
				"Author":  "Author 1",
				"Country": "USA",
			}
			metaJson, _ := json.Marshal(&meta)

			productPatchForm := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			productPatchForm.Meta = metaJson

			productPatchModel, _ := productPatchForm.ToModel()
			productPatchModel.TenantId = 1

			productDto := productPatchModel.ToDto()

			IProductRepo := new(mocks.IProductRepo)
			IProductRepo.On("Patch", mock.Anything, 1).Return(productPatchModel, fmt.Errorf("Some error down call chain"))

			cfg := config.AppConfig()
			lang, err := app.InitLocale(cfg)
			assert.NoError(t, err)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
				Lang:        lang,
			})

			u, _ := ps.Patch(productPatchForm, 1)

			assert.NotEqual(t, u, productDto)
			IProductRepo.AssertExpectations(t)
		})
	})

	t.Run("Get Product Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			product := &model.ProductDetail{
				Name: "product 1",
				Code: "code",
			}

			productDto := product.ToProductDetailDto()

			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("Get", 1, 1).Return(product, nil)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, err := ps.Get(1, 1)

			assert.NoError(t, err)
			assert.Equal(t, u, productDto)
		})
		t.Run("Error", func(t *testing.T) {
			product := &model.ProductDetail{
				Name: "product 1",
				Code: "code",
			}

			productDto := product.ToProductDetailDto()

			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("Get", 1, 1).Return(nil, fmt.Errorf("Some error down call chain"))

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, err := ps.Get(1, 1)

			assert.Error(t, err)
			assert.NotEqual(t, u, productDto)
		})
	})

	t.Run("Add Product Service", func(t *testing.T) {
		t.Run("Addproduct_200_Success", func(t *testing.T) {
			meta := `{"Author":"Author 1","Country":"USA"}`
			productForm := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participant",
				Meta:        datatypes.JSON(meta),
				IsActive:    1,
			}

			productModel, _ := productForm.ToModel()
			productDto := productModel.ToDto()
			//Todo - Get from token
			productModel.TenantId = 1
			IProductRepo := new(mocks.IProductRepo)

			IProductRepo.On("Add", mock.Anything).Return(productModel, nil, nil)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, err := ps.Add(1, productForm)

			assert.NoError(t, err)
			assert.Equal(t, u, productDto)

		})
		t.Run("Addproduct_CheckProductSchema_Meta_Success", func(t *testing.T) {

			meta := map[string]interface{}{
				"$schema":     "http://json-schema.org/draft-04/schema#",
				"description": "product creation",
				"properties": map[string]interface{}{
					"country": map[string]interface{}{
						"description": "Product Country", "minLength": float64(1), "type": "string",
					}, "author": map[string]interface{}{
						"description": "Product Address", "minLength": float64(1), "type": "string",
					},
				},
				"required": []interface{}{},
				"title":    "product",
				"type":     "object",
			}

			metaJson, _ := json.Marshal(&meta)

			IProductRepo := new(mocks.IProductRepo)
			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, _, err := ps.CheckProductSchema(0, metaJson, "validateMetaData")

			assert.NoError(t, err)
			assert.Equal(t, u, meta)

		})
		t.Run("Addproduct_ConvertMapToJson_Success", func(t *testing.T) {

			meta := map[string]interface{}{
				"$schema":     "http://json-schema.org/draft-04/schema#",
				"description": "product creation",
				"properties": map[string]interface{}{
					"country": map[string]interface{}{
						"description": "Product Country", "minLength": float64(1), "type": "string",
					}, "author": map[string]interface{}{
						"description": "Product Address", "minLength": float64(1), "type": "string",
					},
				},
				"required": []interface{}{},
				"title":    "product",
				"type":     "object",
			}

			IProductRepo := new(mocks.IProductRepo)
			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
			})

			u, err := ps.ConvertMapToJson(&meta)

			assert.NoError(t, err)
			assert.NotEqual(t, u, meta)

		})
		t.Run("Addproduct_500_Failure", func(t *testing.T) {
			meta := map[string]interface{}{
				"Author":  "Author 1",
				"Country": "USA",
			}
			metaJson, _ := json.Marshal(&meta)

			productForm := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
				IsActive:    0,
			}
			productForm.Meta = metaJson

			productModel, _ := productForm.ToModel()
			productModel.TenantId = 1

			productDto := productModel.ToDto()

			IProductRepo := new(mocks.IProductRepo)
			IProductRepo.On("Add", productModel).Return(productModel, fmt.Errorf("Some error down call chain"))

			cfg := config.AppConfig()
			lang, err := app.InitLocale(cfg)
			assert.NoError(t, err)

			ps := NewService(ServiceConfig{
				ProductRepo: IProductRepo,
				Lang:        lang,
			})

			u, _ := ps.Add(1, productForm)
			assert.NotEqual(t, u, productDto)
			IProductRepo.AssertExpectations(t)
		})
	})
}
