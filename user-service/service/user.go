package service

import (
	"fmt"
	"micro/app/locale"
	"micro/model"
	"micro/repo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ListUsers(c *gin.Context) (model.Users, error) {

	db := c.MustGet("db").(*gorm.DB)
	lang := c.MustGet("lang").(*locale.Locale)
	log := c.MustGet("log").(*logrus.Logger)

	fmt.Println("Calling List user with lang : ", lang.Get("hi", "Bob"))
	lang.SetLang("en-US")
	fmt.Println("Calling List user with lang : ", lang.Get("hi", "Bob"))

	log.Info("Sample log from listuser")

	users, err := repo.ListUsers(db)

	if err != nil {
		return nil, err
	}

	return users, nil
}
