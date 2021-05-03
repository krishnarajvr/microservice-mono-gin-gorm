package service

import (
	"fmt"
	"micro/app"
	"micro/config"
	"testing"

	"micro/module/client/mocks"
	"micro/module/client/model"

	common "github.com/krishnarajvr/micro-common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestClientCredential(t *testing.T) {
	t.Run("List ClientCredential Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredential{
				Name:     "client 1",
				Type:     "type",
				TenantId: 1,
			}
			page := common.Pagination{
				PageNumber: "ID",
				PageOrder:  "DESC",
				PageOffset: "0",
				Search:     "",
			}
			clients := model.ClientCredential{mockClientCredentialResp}
			clientsDtos := clients.ToCredentialDto()

			IClientRepo := new(mocks.IClientRepo)

			IClientRepo.On("List", page).Return(clients, nil, nil)

			ps := NewService(ServiceConfig{
				ClientRepo: IClientRepo,
			})

			u, _, err := ps.ListCredentials(page)

			assert.NoError(t, err)
			assert.Equal(t, u, clientsDtos)
		})
	})

	t.Run("Patch ClientCredential Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mockClientCredentialResp := &model.ClientCredential{
				Name:        "client 1",
				Type:        "type",
				TenantId:    1,
				Code:        "code",
				ReferenceId: "anna@gmail.com",
			}
			mockClientPatchCredentialReq := &model.ClientCredentialPatchForm{
				Name: "client 1",
				Type: "type",
				Code: "code",
			}

			IClientRepo := new(mocks.IClientRepo)

			IClientRepo.On("Patch", mock.Anything, 1).Return(mockClientCredentialResp, nil, nil)

			ps := NewService(ServiceConfig{
				ClientRepo: IClientRepo,
			})

			_, err := ps.PatchCredential(mockClientPatchCredentialReq, 1)

			assert.NoError(t, err)
		})
	})

	t.Run("Get ClientCredential Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			client := &model.ClientCredential{
				Name:        "client 1",
				Type:        "type",
				TenantId:    1,
				Code:        "code",
				ReferenceId: "bob@gmail.com",
			}

			clientCredntialDto := client.ToCredentialDto()

			IClientRepo := new(mocks.IClientRepo)

			IClientRepo.On("Get", 1).Return(client, nil, nil)

			ps := NewService(ServiceConfig{
				ClientRepo: IClientRepo,
			})

			u, err := ps.GetCredential(1)

			assert.NoError(t, err)
			assert.Equal(t, u, clientCredntialDto)
		})
	})

	t.Run("Add_ClientCredential_Service", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			clientForm := &model.ClientCredentialForm{
				Name:        "client 1",
				Type:        "type",
				Code:        "code",
				ReferenceId: "bob@gmail.com",
				IsActive:    1,
			}

			clientModel, _ := clientForm.ToModel()
			clientCredntialDto := clientModel.ToCredentialDto()
			clientModel.TenantId = 1
			IClientRepo := new(mocks.IClientRepo)

			IClientRepo.On("AddCredential", mock.Anything).Return(clientModel, nil, nil)

			ps := NewService(ServiceConfig{
				ClientRepo: IClientRepo,
			})

			u, err := ps.AddCredential(clientForm, 1)

			assert.NoError(t, err)
			assert.Equal(t, u, clientCredntialDto)

		})
		t.Run("Failure", func(t *testing.T) {
			clientForm := &model.ClientCredentialForm{
				Name:        "client 1",
				Type:        "type",
				Code:        "code",
				ReferenceId: "bob@gmail.com",
				IsActive:    1,
			}

			clientModel, _ := clientForm.ToModel()
			clientCredntialDto := clientModel.ToCredentialDto()
			clientModel.TenantId = 1
			IClientRepo := new(mocks.IClientRepo)

			IClientRepo.On("AddCredential", mock.Anything).Return(clientModel, fmt.Errorf("Some error down call chain"))

			cfg := config.AppConfig()
			lang, err := app.InitLocale(cfg)

			assert.NoError(t, err)

			ps := NewService(ServiceConfig{
				ClientRepo: IClientRepo,
				Lang:       lang,
			})

			u, _ := ps.AddCredential(clientForm, 1)
			assert.NotEqual(t, u, clientCredntialDto)
			IClientRepo.AssertExpectations(t)
		})
	})
}
