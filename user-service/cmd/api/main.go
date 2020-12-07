package main

import (
	"micro/app/database"
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
	dbs, _ := database.Initialize()

	app := gin.Default() // create gin app
	app.Use(database.Inject(dbs))
	route.Apply(app)                   // apply api router
	app.Run(":" + appConf.Server.Port) // listen to given port
}
