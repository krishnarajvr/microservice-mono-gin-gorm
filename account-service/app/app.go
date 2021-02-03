package app

import (
	"micro/app/locale"
	"micro/config"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/middleware"
)

//AppConfig - Application config
type AppConfig struct {
	Dbs     *Dbs
	Lang    *locale.Locale
	Router  *gin.Engine
	BaseURL string
}

// InitRouter - Create gin router
func InitRouter() (*gin.Engine, error) {
	cfg := config.AppConfig()
	router := gin.Default()
	router.Use(middleware.LoggerToFile(cfg.Log.LogFilePath, cfg.Log.LogFileName))
	router.Use(gin.Recovery())

	return router, nil
}
