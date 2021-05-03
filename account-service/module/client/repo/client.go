package repo

import (
	"micro/module/client/model"
	"strings"

	"gorm.io/gorm"

	common "github.com/krishnarajvr/micro-common"
)

type IClientRepo interface {
	ListCredentials(page common.Pagination) (model.ClientCredentials, *common.PageResult, error)
	AddCredential(form *model.ClientCredential) (*model.ClientCredential, error)
	GetCredential(id int) (*model.ClientCredential, error)
	PatchCredential(form *model.ClientCredentialPatchForm, id int) (*model.ClientCredential, error)
	CheckCredentials(form *model.ClientLoginForm) (*model.ClientCredential, error)
	GetClientCredentialRoles(clientCredentialId int) ([]string, error)
}

type ClientRepo struct {
	DB *gorm.DB
}

func NewClientRepo(db *gorm.DB) ClientRepo {
	return ClientRepo{
		DB: db,
	}
}

func (r ClientRepo) ListCredentials(page common.Pagination) (model.ClientCredentials, *common.PageResult, error) {
	clientCredentials := make([]*model.ClientCredential, 0)
	var totalCount int64

	if err := r.DB.Table("client_credentials").Count(&totalCount).Error; err != nil {
		return nil, nil, err
	}

	if err := r.DB.Scopes(common.Paginate(page)).Find(&clientCredentials).Error; err != nil {
		return nil, nil, err
	}

	pageResult := common.PageInfo(page, totalCount)

	if len(clientCredentials) == 0 {
		return nil, nil, nil
	}

	return clientCredentials, &pageResult, nil
}

func (r ClientRepo) AddCredential(clientCredential *model.ClientCredential) (*model.ClientCredential, error) {
	var id uint
	row := r.DB.Table("roles").
		Select("id").
		Where("code = ?", strings.ToUpper(clientCredential.Type)).
		Row()
	row.Scan(&id)

	errFromDB := r.DB.Transaction(func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Error; err != nil {
			return err
		}

		if err := tx.Create(&clientCredential).Error; err != nil {
			tx.Rollback()
			return err
		}

		ccr := model.ClientCredentialRole{ClientCredentialID: clientCredential.ID, RoleID: id}

		if err := tx.Create(&ccr).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Error
	})

	if errFromDB != nil {
		return nil, errFromDB
	}

	return clientCredential, nil
}

func (r ClientRepo) GetCredential(id int) (*model.ClientCredential, error) {
	clientCredential := new(model.ClientCredential)

	if err := r.DB.Where("id = ?", id).First(&clientCredential).Error; err != nil {
		return nil, err
	}

	return clientCredential, nil
}

func (r ClientRepo) PatchCredential(form *model.ClientCredentialPatchForm, id int) (*model.ClientCredential, error) {
	clientCredential, err := form.ToModel()

	if err != nil {
		return nil, err
	}

	if err := r.DB.Where("id = ?", id).Updates(&clientCredential).Error; err != nil {
		return nil, err
	}

	return clientCredential, nil
}

func (r ClientRepo) CheckCredentials(form *model.ClientLoginForm) (*model.ClientCredential, error) {
	clientCredential := new(model.ClientCredential)
	err := r.DB.Table("client_credentials").
		Where("code = ?", form.ClientId).
		Where("secret = ?", form.ClientSecret).
		First(&clientCredential).Error

	if err != nil {
		return nil, err
	}

	return clientCredential, nil
}

func (r ClientRepo) GetClientCredentialRoles(clientCredentialId int) ([]string, error) {
	var roles []string
	err := r.DB.Table(`roles`).
		Select(`LOWER(roles.code) as code`).
		Joins(`JOIN client_credential_roles on client_credential_roles.role_id = roles.id`).
		Where(`client_credential_roles.client_credential_id = ? `, clientCredentialId).
		Find(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}
