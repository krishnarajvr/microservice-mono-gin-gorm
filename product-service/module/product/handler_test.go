package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro/app"
	"micro/module/product/mocks"
	"micro/module/product/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/datatypes"

	common "github.com/krishnarajvr/micro-common"

	"github.com/krishnarajvr/micro-common/middleware"

	"micro/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProduct(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	t.Run("Patch product", func(t *testing.T) {

		t.Run("ProductPatch_Meta_200_Success", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				ID:          1,
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			addform := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			addform.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)

			mockProductService.On("Patch", form, 1).Return(mockProductResp, nil)
			mockProductService.On("Add", 0, addform).Return(mockProductResp, nil)
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, nil)
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, nil)

			mockGetProductResp := &model.ProductDetailDto{
				ID:   1,
				Name: "john",
				Code: "1001",
			}
			mockProductService.On("Get", 0, 1).Return(mockGetProductResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusOK, rr.Code)
		})
		t.Run("ProductPatch_Meta_ConvertMapToJson_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				ID:          1,
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)

			mockProductService.On("Patch", form, 1).Return(mockProductResp, nil)
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, nil)
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, fmt.Errorf("Some error down call chain"))

			mockGetProductResp := &model.ProductDetailDto{
				ID:   1,
				Name: "john",
				Code: "1001",
			}
			mockProductService.On("Get", 0, 1).Return(mockGetProductResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductPatch_Meta_CheckProductSchema_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				ID:          1,
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)

			mockProductService.On("Patch", form, 1).Return(mockProductResp, nil)
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, fmt.Errorf("Some error down call chain"))
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, nil)

			mockGetProductResp := &model.ProductDetailDto{
				ID:   1,
				Name: "john",
				Code: "1001",
			}
			mockProductService.On("Get", 0, 1).Return(mockGetProductResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductPatch_Patch_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				ID:          1,
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)

			mockProductService.On("Patch", form, 1).Return(mockProductResp, fmt.Errorf("Some error down call chain"))
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, nil)
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, nil)

			mockGetProductResp := &model.ProductDetailDto{
				ID:   1,
				Name: "john",
				Code: "1001",
			}
			mockProductService.On("Get", 0, 1).Return(mockGetProductResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductPatch_ValidateId_500_Failure", func(t *testing.T) {

			mockProductService := new(mocks.IProductService)
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/wewewe", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductPatch_Get_500_Failure", func(t *testing.T) {
			mockProductResp := &model.ProductDto{
				ID:          1,
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductPatchForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}

			address := map[string]interface{}{
				"zipCode": "636805",
			}

			assert.NoError(t, err)
			addressMockProduct := map[string]interface{}{
				"zipCode": "636805",
			}
			addressMockProductJson, err := json.Marshal(&addressMockProduct)
			assert.NoError(t, err)

			assert.NoError(t, err)

			mockProductService := new(mocks.IProductService)

			mockProductService.On("Patch", form, 1).Return(mockProductResp, nil)
			mockProductService.On("ConvertMapToJson", &addressMockProduct).Return(addressDataTypeJson, nil)

			mockGetProductResp := &model.ProductDetailDto{
				ID:   1,
				Name: "john",
				Code: "1001",
			}
			mockProductService.On("Get", 0, 1).Return(mockGetProductResp, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductPatch_404_NotFound", func(t *testing.T) {

			mockProductService := new(mocks.IProductService)
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"Name":        "john",
				"Code":        "1001",
				"Description": "john has 3 participants",
				"Meta":        "{\"author\":\"USA\",\"country\":\"USA\"}",
				"IsActive":    0,
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPatch, "/productes/1", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusNotFound, rr.Code)
		})

	})

	t.Run("Add product", func(t *testing.T) {

		t.Run("ProductAdd_Validation_400_Badrequest", func(t *testing.T) {

			mockProductService := new(mocks.IProductService)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"code":        "1001",
				"description": "john has 3 participants",
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
		t.Run("ProductAdd_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
				IsActive:    0,
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)
			mockProductService.On("Add", 0, form).Return(mockProductResp, fmt.Errorf("Some error down call chain"))
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, nil)
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductAdd_Meta_CheckProductSchema_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
				IsActive:    0,
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)
			mockProductService.On("Add", 0, form).Return(mockProductResp, nil)
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, fmt.Errorf("Some error down call chain"))
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductAdd_Meta_ConvertMapToJson_500_Failure", func(t *testing.T) {

			mockProductResp := &model.ProductDto{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
			}
			form := &model.ProductForm{
				Name:        "john",
				Code:        "1001",
				Description: "john has 3 participants",
				IsActive:    0,
			}

			meta := map[string]interface{}{
				"author": "Author 1",
			}

			metaJson, err := json.Marshal(&meta)
			form.Meta = metaJson
			assert.NoError(t, err)

			metaMockProduct := map[string]interface{}{
				"author": "Author 1",
			}
			metaMockProductJson, err := json.Marshal(&metaMockProduct)
			mockProductResp.Meta = metaMockProductJson
			assert.NoError(t, err)

			var metaSchJson datatypes.JSON
			metaSchJson = []byte(`{
				"$schema": "http://json-schema.org/draft-04/schema#",
				"title": "product",
				"description": "product creation",
				"type": "object",
				"properties": {
					"author": {
						"description": "Product Address",
						"type": "string",
						"minLength": 1
					},
					"country": {
						"description": "Product Country",
						"type": "string",
						"minLength": 1
					}
				},
				"required": []
			}`)

			var metaSchemaDef map[string]interface{}
			err = json.Unmarshal([]byte(metaSchJson), &metaSchemaDef)

			assert.NoError(t, err)

			var metaDataTypeJson datatypes.JSON
			metaDataTypeJson = []byte(`{"author":"USA"}`)
			var metaData = `{"author":"USA"}`

			var formMeta *map[string]interface{}
			json.Unmarshal([]byte(metaData), &formMeta)
			mockProductService := new(mocks.IProductService)
			mockProductService.On("Add", 0, form).Return(mockProductResp, nil)
			mockProductService.On("CheckProductSchema", mock.Anything, metaJson, "validateMetaData").Return(metaSchemaDef, formMeta, nil)
			mockProductService.On("ConvertMapToJson", &metaMockProduct).Return(metaDataTypeJson, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			// don't need a middleware as we don't yet have authorized Product
			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			reqBody, err := json.Marshal(gin.H{
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"meta":        meta,
				"is_active":   0,
			})

			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)

			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
		t.Run("ProductAdd_404_NotFound", func(t *testing.T) {

			mockproductService := new(mocks.IProductService)
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			lang, _ := app.InitLocale(cfg)

			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockproductService,
				Lang:           lang,
			})

			reqBody, err := json.Marshal(gin.H{
				"name":        "john",
				"code":        "1001",
				"description": "john has 3 participants",
				"is_active":   0,
			})
			assert.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/productes", bytes.NewReader(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)
			router.ServeHTTP(rr, request)

			fmt.Println(rr.Code)
			assert.Equal(t, http.StatusNotFound, rr.Code)
		})
	})

	t.Run("List_Product", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockProductResp := &model.ProductList{
				Name: "product 1",
				ID:   1,
				Code: "code",
			}

			products := model.ProductsList{mockProductResp}
			productsDtos := products.ToDto()

			mockProductService := new(mocks.IProductService)

			page := common.Pagination{
				PageNumber: "1",
				PageOffset: "",
				PageSize:   "25",
				PageOrder:  "",
				Search:     "",
			}

			var mockFilters model.ProductFilterList
			mockFilters.Code = ""
			mockFilters.Name = ""

			mockProductService.On("List", 0, page, mockFilters).Return(productsDtos, nil, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products", bytes.NewBuffer(nil))
			request.Header.Set("X-Tenant-Id", mock.Anything)

			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			router.ServeHTTP(rr, request)
			assert.Equal(t, 200, rr.Code)
		})
		t.Run("500_InternalServerError", func(t *testing.T) {
			mockProductService := new(mocks.IProductService)
			page := common.Pagination{
				PageNumber: "1",
				PageOrder:  "",
				PageOffset: "",
				Search:     "",
				PageSize:   "25",
			}

			var mockFilters model.ProductFilterList
			mockFilters.Code = ""
			mockFilters.Name = ""

			mockProductService.On("List", 0, page, mockFilters).Return(nil, nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products", nil)
			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			router.ServeHTTP(rr, request)
			fmt.Println(rr)
			assert.Equal(t, 500, rr.Code)
		})
	})

	t.Run("Get_Product", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockProductResp := &model.ProductDetail{
				Name:  "product 1",
				ID:    1,
				Code:  "code",
				Email: "bob@gmail.com",
			}

			productDto := mockProductResp.ToProductDetailDto()
			mockProductService := new(mocks.IProductService)
			mockProductService.On("Get", 0, 1).Return(productDto, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products/1", nil)
			request.Header.Set("X-Tenant-Id", mock.Anything)

			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")

			router.ServeHTTP(rr, request)
			assert.Equal(t, 200, rr.Code)
		})
		t.Run("500_InternalServerError", func(t *testing.T) {

			mockProductService := new(mocks.IProductService)

			mockProductService.On("Get", 0, 1).Return(nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products/1", nil)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			assert.NoError(t, err)
			router.ServeHTTP(rr, request)
			assert.Equal(t, 500, rr.Code)
		})
		t.Run("InvalidId_Error", func(t *testing.T) {

			mockProductService := new(mocks.IProductService)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products/abckd", nil)
			request.Header.Set("Product-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)

			assert.NoError(t, err)
			router.ServeHTTP(rr, request)
			assert.Equal(t, 500, rr.Code)
		})
		t.Run("ProductId_CheckError", func(t *testing.T) {
			mockProductResp := &model.ProductDetail{
				Name:  "product 1",
				ID:    0,
				Code:  "code",
				Email: "bob@gmail.com",
			}

			productDto := mockProductResp.ToProductDetailDto()
			mockProductService := new(mocks.IProductService)
			mockProductService.On("Get", 0, 1).Return(productDto, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			excludeList := map[string]interface{}{}
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:              router,
				ProductService: mockProductService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/products/1", nil)
			request.Header.Set("X-Tenant-Id", mock.Anything)

			assert.NoError(t, err)
			request.Header.Set("Product-Type", "application/json")

			router.ServeHTTP(rr, request)
			assert.Equal(t, 500, rr.Code)
		})
	})
}
