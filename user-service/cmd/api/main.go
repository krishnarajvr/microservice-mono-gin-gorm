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
