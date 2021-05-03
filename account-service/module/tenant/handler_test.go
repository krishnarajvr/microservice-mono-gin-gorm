package tenant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro/module/tenant/mocks"
	"micro/module/tenant/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"micro/config"

	"github.com/gin-gonic/gin"
	common "github.com/krishnarajvr/micro-common"
	"github.com/krishnarajvr/micro-common/middleware"
	"github.com/stretchr/testify/assert"
)

func TestTenant(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("List tenant", func(t *testing.T) {

		mockTenantResp := &model.TenantDto{
			ID:   1,
			Name: "Tenant 1",
			Code: "Code 1",
		}
		dtos := model.TenantDtos{mockTenantResp}

		mockTenantService := new(mocks.ITenantService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockTenantService.On("List", page).Return(dtos, nil, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// don't need a middleware as we don't yet have authorized tenant
		router := gin.Default()
		cfg := config.AppConfig()
		router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

		InitRoutes(HandlerConfig{
			R:             router,
			TenantService: mockTenantService,
		})

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"tenant": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/tenants", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

	t.Run("List tenant Error", func(t *testing.T) {

		mockTenantService := new(mocks.ITenantService)

		page := common.Pagination{
			Sort:   "ID",
			Order:  "DESC",
			Offset: "0",
			Limit:  "25",
			Search: "",
		}

		mockTenantService.On("List", page).Return(nil, nil, fmt.Errorf("Some error down call chain"))

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		cfg := config.AppConfig()
		router := gin.Default()
		router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

		InitRoutes(HandlerConfig{
			R:             router,
			TenantService: mockTenantService,
		})

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodGet, "/tenants", bytes.NewBuffer(nil))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		fmt.Println(rr)

		assert.Equal(t, 200, rr.Code)

	})

}
