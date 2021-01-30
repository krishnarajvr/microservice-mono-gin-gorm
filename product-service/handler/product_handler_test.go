package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro/model"
	"micro/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("List product", func(t *testing.T) {

		mockProductResp := &model.ProductDto{
			ID:          1,
			Name:        "Product 1",
			Code:        "Code 1",
			Description: "Description 1",
		}
		dtos := model.ProductDtos{mockProductResp}

		mockProductService := new(mocks.IProductService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockProductService.On("List", page).Return(dtos, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized product
		router := gin.Default()

		NewHandler(&Config{
			R:              router,
			ProductService: mockProductService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"product": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/products", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

	t.Run("List product Error", func(t *testing.T) {

		mockProductService := new(mocks.IProductService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockProductService.On("List", page).Return(nil, fmt.Errorf("Some error down call chain"))

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized product
		router := gin.Default()

		NewHandler(&Config{
			R:              router,
			ProductService: mockProductService,
		})

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/products", bytes.NewBuffer(nil))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

}
