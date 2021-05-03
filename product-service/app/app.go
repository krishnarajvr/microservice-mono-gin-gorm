package app

import (
	"log"
	lgorm "micro/app/database/gorm"
	"micro/config"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/locale"
	"github.com/krishnarajvr/micro-common/middleware"
)

//AppConfig - Application config
type AppConfig struct {
	Dbs     *Dbs
	Lang    *locale.Locale
	Router  *gin.Engine
	BaseURL string
	Cfg     *config.Conf
}

type Dbs struct {
	DB *gorm.DB
}

// InitRouter - Create gin router
func InitRouter(cfg *config.Conf, excludeList map[string]interface{}) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
	router.Use(middleware.TenantValidator(excludeList))
	router.Use(gin.Recovery())

	return router, nil
}

// InitLocale - Create locale object
func InitLocale(cfg *config.Conf) (*locale.Locale, error) {
	langLocale := locale.Locale{}
	dir, err := os.Getwd()

	if err != nil {
		log.Print("Not able to get current working director")
		panic(err)
	}

	log.Println("Locale path: " + dir + "/locale/*/*")
	lang := langLocale.New(cfg.App.Lang, dir+"/locale/*/*", "en-GR", "en-US", "zh-CN")

	return lang, nil
}

// InitDS establishes connections to fields in dataSources
func InitDS(config *config.Conf) (*Dbs, error) {
	db, err := lgorm.New(config)

	if err != nil {
		log.Println("Connection failed")
		panic(err)
	}

	return &Dbs{
		DB: db,
	}, nil
}

//Close to be used in graceful server shutdown
func Close(d *Dbs) error {
	//Todo Check the error
	//return d.DB.Close()
	return nil
}
