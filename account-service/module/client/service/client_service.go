package service

import (
	"encoding/json"
	"micro/module/client/model"
	"micro/module/client/repo"
	"micro/util/token"

	common "github.com/krishnarajvr/micro-common"

	"github.com/google/uuid"
	"github.com/krishnarajvr/micro-common/locale"
)

type IClientService interface {
	ListCredentials(page common.Pagination) (model.ClientCredentialDtos, *common.PageResult, error)
	AddCredential(content *model.ClientCredentialForm, tenantId int) (*model.ClientCredentialDto, error)
	GetCredential(id int) (*model.ClientCredentialDto, error)
	PatchCredential(form *model.ClientCredentialPatchForm, id int) (*model.ClientCredentialDto, error)
	CheckCredentials(form *model.ClientLoginForm) (*model.ClientCredential, error)
	GetClientToken(credential *model.ClientCredential, roles []string) (*token.TokenDetails, error)
	GetClientCredentialRoles(clientCredentialId int) ([]string, error)
	RefreshToken(refreshToken string) (*token.TokenDetails, error)
}

type ServiceConfig struct {
	ClientRepo repo.IClientRepo
	Lang       *locale.Locale
	Token      *token.Token
}

type Service struct {
	ClientRepo repo.IClientRepo
	Lang       *locale.Locale
	Token      *token.Token
}

func NewService(c ServiceConfig) IClientService {
	return &Service{
		ClientRepo: c.ClientRepo,
		Lang:       c.Lang,
		Token:      c.Token,
	}
}

func (s *Service) ListCredentials(page common.Pagination) (model.ClientCredentialDtos, *common.PageResult, error) {
	clientCredentials, pageResult, err := s.ClientRepo.ListCredentials(page)

	if err != nil {
		return nil, nil, err
	}

	clientCredentialsDtos := clientCredentials.ToCredentialDtos()

	return clientCredentialsDtos, pageResult, err
}

func (s *Service) AddCredential(form *model.ClientCredentialForm, tenantId int) (*model.ClientCredentialDto, error) {
	clientCredentialModel, err := form.ToModel()
	//Todo - Get from token
	clientCredentialModel.TenantId = tenantId
	clientCredentialModel.Secret = uuid.New().String()

	if err != nil {
		return nil, err
	}

	clientCredential, err := s.ClientRepo.AddCredential(clientCredentialModel)
	if err != nil {
		return nil, err
	}

	clientCredntialDto := clientCredential.ToCredentialDto()

	return clientCredntialDto, nil
}

func (s *Service) GetCredential(id int) (*model.ClientCredentialDto, error) {
	clientCredential, err := s.ClientRepo.GetCredential(id)

	if err != nil {
		return nil, err
	}

	clientCredntialDto := clientCredential.ToCredentialDto()

	return clientCredntialDto, nil
}

func (s *Service) PatchCredential(form *model.ClientCredentialPatchForm, id int) (*model.ClientCredentialDto, error) {
	clientCredential, err := s.ClientRepo.PatchCredential(form, id)

	if err != nil {
		return nil, err
	}

	clientCredntialDto := clientCredential.ToCredentialDto()

	return clientCredntialDto, nil
}

func (s *Service) CheckCredentials(form *model.ClientLoginForm) (*model.ClientCredential, error) {
	clientCredential, err := s.ClientRepo.CheckCredentials(form)

	if err != nil {
		return nil, err
	}

	return clientCredential, nil
}

func (s *Service) GetClientToken(credential *model.ClientCredential, roles []string) (*token.TokenDetails, error) {
	subject := "Vendor"
	var payload map[string]interface{}

	if len(credential.Payload) != 0 {
		json.Unmarshal(credential.Payload, &payload)

		if payload["name"] != nil {
			subject = payload["name"].(string)
		}
	}

	tokenData := token.TokenData{
		ReferenceId: credential.ReferenceId,
		Type:        credential.Type,
		TenantId:    int64(credential.TenantId),
		Subject:     subject,
		Admin:       false,
		Roles:       roles,
	}

	token, _ := s.Token.CreateToken(&tokenData)

	return token, nil
}

func (s *Service) GetClientCredentialRoles(clientCredentialId int) ([]string, error) {
	roles, err := s.ClientRepo.GetClientCredentialRoles(clientCredentialId)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *Service) RefreshToken(refreshToken string) (*token.TokenDetails, error) {
	token, err := s.Token.Refresh(refreshToken, "vendor", "ThirdParty")

	if err != nil {
		return nil, err
	}

	return token, nil
}
