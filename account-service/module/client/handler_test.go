package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"micro/module/client/mocks"
	"micro/module/client/model"
	"net/http"
	"net/http/httptest"
	"testing"

	common "github.com/krishnarajvr/micro-common"

	"micro/config"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientCredential(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	t.Run("Patch_ClientCredential_handler", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				ID:          1,
				Name:        "client 1",
				Type:        "logPath",
				Code:        "code",
				ReferenceId: "mailid@gmail.com",
			}

			mockClientService := new(mocks.IClientService)

			mockClientService.On("Patch", mock.Anything, 1).Return(mockClientCredentialResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"id":          1,
				"name":        "abc",
				"code":        "test123",
				"secret":      "address",
				"referenceId": "bob@yes.com",
				"credentails": "{\"key\": \"value\"}",
				"type":        "logPath",
				"isActive":    true,
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPatch, "/clients/1", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusOK, rr.Code)
		})

		t.Run("500_InternalServerError", func(t *testing.T) {
			mockClientService := new(mocks.IClientService)

			mockClientService.On("Patch", mock.Anything, 1).Return(nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"id":             1,
				"clientId":       0,
				"name":           "ABC",
				"code":           "",
				"contact_detail": "",
				"contact_email":  "a@abc.com",
				"credentails":    "{\"a\": 1}",
				"logo_path":      "asdasd",
				"createdAt":      "1999-01-02T12:20:20Z",
				"updatedAt":      "1999-01-02T12:20:20Z",
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPatch, "/clients/1", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	})

	t.Run("Add_ClientCredential_handler", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				ID:          1,
				Name:        "client 1",
				Code:        "code",
				ReferenceId: "1234",
				Type:        "vendor",
				IsActive:    1,
			}

			form := &model.ClientCredentialForm{
				Name:        "client 1",
				Code:        "code",
				ReferenceId: "1234",
				Type:        "vendor",
				IsActive:    1,
			}

			mockClientService := new(mocks.IClientService)

			mockClientService.On("AddCredential", form, mock.Anything).Return(mockClientCredentialResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"name":        "client 1",
				"code":        "code",
				"referenceId": "1234",
				"type":        "vendor",
				"isActive":    1,
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPost, "/clientCredentials", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("X-Tenant-Id", mock.Anything)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusOK, rr.Code)
		})

		t.Run("400_BadRequest", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				Name:        "client 1",
				Type:        "vendor",
				Code:        "code",
				ReferenceId: "1234",
			}

			form := &model.ClientCredentialForm{
				Type:        "vendor",
				Code:        "code",
				ReferenceId: "1234",
			}
			mockClientService := new(mocks.IClientService)

			mockClientService.On("AddCredential", form, mock.Anything).Return(mockClientCredentialResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"Type":        "vendor",
				"Code":        "code",
				"ReferenceId": "1234",
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPost, "/clientCredentials", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("X-Tenant-Id", mock.Anything)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})

		t.Run("403_Forbidden", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				Name:        "client 1",
				Type:        "vendor",
				Code:        "code",
				ReferenceId: "1234",
			}

			form := &model.ClientCredentialForm{
				Type:        "vendor",
				Code:        "code",
				ReferenceId: "1234",
			}
			mockClientService := new(mocks.IClientService)

			mockClientService.On("AddCredential", form, mock.Anything).Return(mockClientCredentialResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			excludeList := map[string]interface{}{}
			router := gin.Default()
			cfg := config.AppConfig()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
			router.Use(middleware.TenantValidator(excludeList))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"Type":        "vendor",
				"Code":        "code",
				"ReferenceId": "1234",
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPost, "/clientCredentials", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusForbidden, rr.Code)
		})

		t.Run("500_InternalServerError", func(t *testing.T) {
			mockClientService := new(mocks.IClientService)

			mockClientService.On("AddCredential", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// create a request body with valid data
			reqBody, err := json.Marshal(gin.H{
				"name":        "gandhi",
				"code":        "mahetma",
				"secret":      "gujar",
				"referenceId": "1234",
				"type":        "logPath",
			})

			assert.NoError(t, err)

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodPost, "/clientCredentials", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Tenant-Id", mock.Anything)
			router.ServeHTTP(rr, request)
			assert.Equal(t, http.StatusInternalServerError, rr.Code)
		})
	})

	t.Run("List_ClientCredential_handler", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				Name:        "client 1",
				Type:        "logPath",
				ID:          1,
				Code:        "code",
				ReferenceId: "Doe@gmail.com",
			}

			dtos := model.ClientCredentialDtos{mockClientCredentialResp}
			mockClientService := new(mocks.IClientService)

			page := common.Pagination{
				PageNumber: "DESC",
				PageOffset: "",
				PageSize:   "25",
				PageOrder:  "",
				Search:     "",
			}

			mockClientService.On("List", page).Return(dtos, nil, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/clients", bytes.NewBuffer(nil))
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)
			assert.Equal(t, 200, rr.Code)
		})

		t.Run("500_InternalServerError", func(t *testing.T) {
			mockClientService := new(mocks.IClientService)
			page := common.Pagination{
				PageNumber: "ID",
				PageOrder:  "DESC",
				PageOffset: "0",
				Search:     "",
			}

			mockClientService.On("List", page, "500").Return(nil, nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/clients", nil)
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rr, request)
			fmt.Println(rr)
			assert.Equal(t, 500, rr.Code)
		})
	})

	t.Run("Get_ClientCredential_handler", func(t *testing.T) {
		t.Run("200_Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredentialDto{
				Name:        "client 1",
				ID:          1,
				Type:        "logPath",
				Code:        "code",
				ReferenceId: "1234",
			}

			mockClientService := new(mocks.IClientService)
			mockClientService.On("Get", 1).Return(mockClientCredentialResp, nil)

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()

			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/clients/1", nil)
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(rr, request)
			assert.Equal(t, 200, rr.Code)
		})

		t.Run("500_InternalServerError", func(t *testing.T) {
			mockClientService := new(mocks.IClientService)

			mockClientService.On("Get", 1).Return(nil, nil, fmt.Errorf("Some error down call chain"))

			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			cfg := config.AppConfig()
			router := gin.Default()
			router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))

			InitRoutes(HandlerConfig{
				R:             router,
				ClientService: mockClientService,
			})

			// use bytes.NewBuffer to create a reader
			request, err := http.NewRequest(http.MethodGet, "/clients/1", nil)
			assert.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(rr, request)
			fmt.Println(rr)
			assert.Equal(t, 500, rr.Code)
		})
	})
}
