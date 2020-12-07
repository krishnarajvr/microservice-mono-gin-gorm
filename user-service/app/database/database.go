package database

import (
	"fmt"

	lgorm "micro/app/database/adapter/gorm"
	"micro/config"

	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
	"gorm.io/gorm"
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	appConf := config.AppConfig()
	dbr, err := lgorm.New(appConf)

	if err != nil {
		fmt.Println("Connection failed")
		fmt.Println(appConf)
		panic(err)
	}
	fmt.Println("Connected to database")
	return dbr, err
}
