package service

import (
	"micro/model"
	"micro/repo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUsers(c *gin.Context) (model.Users, error) {

	db := c.MustGet("db").(*gorm.DB)

	users, err := repo.ListUsers(db)

	if err != nil {
		return nil, err
	}

	return users, nil
}
