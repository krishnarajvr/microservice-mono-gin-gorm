package main

import (
	"log"

	"micro/app"
	"micro/common"
	"micro/config"
)

func init() {
	common.LoadEnv()
}

// @title ProductManagement Service API Document
// @version 1.0
// @description List APIs of ProductManagement Service
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

	//Add dependency injection
	ginApp, err := app.Inject(ds)

	ginApp.Run(":" + appConf.Server.Port)
}
