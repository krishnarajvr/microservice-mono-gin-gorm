package app

import (
	"fmt"
	lgorm "micro/app/database/gorm"
	"micro/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type dbs struct {
	DB *gorm.DB
}

// InitDS establishes connections to fields in dataSources
func InitDS() (*dbs, error) {
	appConf := config.AppConfig()
	db, err := lgorm.New(appConf)

	if err != nil {
		fmt.Println("Connection failed")
		fmt.Println(appConf)
		panic(err)
	}

	return &dbs{
		DB: db,
	}, nil
}

// close to be used in graceful server shutdown
func (d *dbs) Close() error {
	return nil
}
