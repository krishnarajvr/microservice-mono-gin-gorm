package repo

import (
	"gorm.io/gorm"

	"micro/model"
)

type IUserRepo interface {
	ListUsers() (model.Users, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (r UserRepo) ListUsers() (model.Users, error) {
	users := make([]*model.User, 0)
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}
