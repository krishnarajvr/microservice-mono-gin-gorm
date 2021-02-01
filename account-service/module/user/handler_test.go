package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro/module/user/mocks"
	"micro/module/user/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("List user", func(t *testing.T) {

		mockUserResp := &model.UserDto{
			ID:   1,
			Name: "User 1",
			Code: "Code 1",
		}
		dtos := model.UserDtos{mockUserResp}

		mockUserService := new(mocks.IUserService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockUserService.On("List", page).Return(dtos, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		InitRoutes(HandlerConfig{
			R:           router,
			UserService: mockUserService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"user": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/users", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

	t.Run("List user Error", func(t *testing.T) {

		mockUserService := new(mocks.IUserService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockUserService.On("List", page).Return(nil, fmt.Errorf("Some error down call chain"))

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		InitRoutes(HandlerConfig{
			R:           router,
			UserService: mockUserService,
		})

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/users", bytes.NewBuffer(nil))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

}
