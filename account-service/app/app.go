package app

import (
	"log"
	"micro/app/locale"
	"micro/config"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/middleware"
)

type AppConfig struct {
	Dbs     *Dbs
	Lang    *locale.Locale
	Router  *gin.Engine
	BaseURL string
}

// Inject all the app dependencies
func InitRouter() (*gin.Engine, error) {
	log.Println("Injecting common classes")

	//Initialize all common classes
	cfg := config.AppConfig()
	// initialize gin.Engine
	router := gin.Default()
	router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
	router.Use(gin.Recovery())
	return router, nil
}
