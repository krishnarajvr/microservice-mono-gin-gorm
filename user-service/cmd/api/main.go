package main

import (
	"log"

	"micro/app"
	"micro/app/locale"
	"micro/app/middleware"

	"micro/config"

	common "github.com/krishnarajvr/micro-common"

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
	log.Println("Starting server...")

	// initialize data sources
	ds, err := app.InitDS()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	ginApp, err := app.Inject(ds)

	lang := locale.Locale{}

	ginApp.Use(middleware.LoggerToFile())
	ginApp.Use(gin.Recovery())
	//app.Use(database.Inject(db))
	ginApp.Use(locale.Inject(lang.New("el-GR")))
	//route.Apply(app)                   // apply api router
	ginApp.Run(":" + appConf.Server.Port) // listen to given port
}
