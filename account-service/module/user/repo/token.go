package repo

import (
	"micro/module/user/model"

	"gorm.io/gorm"
)

type ITokenRepo interface {
	Get(tenantId int, id int) (*model.Token, error)
}

type TokenRepo struct {
	DB *gorm.DB
}

func NewTokenRepo(db *gorm.DB) TokenRepo {
	return TokenRepo{
		DB: db,
	}
}

func (r TokenRepo) Get(tenantId int, id int) (*model.Token, error) {
	token := new(model.Token)
	err := r.DB.
		Where("id = ?", id).
		Where("tenant_id = ? ", tenantId).
		First(&token).Error

	if err != nil {
		return nil, err
	}

	return token, nil
}
