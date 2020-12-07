package repo

import (
	"gorm.io/gorm"

	"micro/model"
)

func ListUsers(db *gorm.DB) (model.Users, error) {
	
	users := make([]*model.User, 0)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}
