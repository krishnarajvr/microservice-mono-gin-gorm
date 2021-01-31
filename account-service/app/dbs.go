package app

import (
	"fmt"
	lgorm "micro/app/database/gorm"
	"micro/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Dbs struct {
	DB *gorm.DB
}

// InitDS establishes connections to fields in dataSources
func InitDS() (*Dbs, error) {
	appConf := config.AppConfig()
	db, err := lgorm.New(appConf)

	if err != nil {
		fmt.Println("Connection failed")
		fmt.Println(appConf)
		panic(err)
	}

	return &Dbs{
		DB: db,
	}, nil
}

// close to be used in graceful server shutdown
func (d *Dbs) Close() error {
	return nil
}
