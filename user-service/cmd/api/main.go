package main

import (
	"micro/app/database"
	"micro/app/locale"
	"micro/app/middleware"
	"micro/common"
	"micro/config"
	"micro/route"

	"github.com/gin-gonic/gin"
)

func init() {
	common.LoadEnv()
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host 127.0.0.1:8081
// @BasePath /v1
func main() {

	appConf := config.AppConfig()
	// initializes database
	db, _ := database.Initialize()
	lang := locale.Locale{}
	app := gin.Default() // create gin app
	app.Use(middleware.LoggerToFile())
	app.Use(gin.Recovery())
	app.Use(database.Inject(db))
	app.Use(locale.Inject(lang.New("el-GR")))
	route.Apply(app)                   // apply api router
	app.Run(":" + appConf.Server.Port) // listen to given port
}
