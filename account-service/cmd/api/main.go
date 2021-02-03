package main

import (
	"log"

	"micro/app"
	"micro/app/locale"
	"micro/config"
	"micro/module"
)

// @title Account Service API Document
// @version 1.0
// @description List APIs of Account Service
// @termsOfService http://swagger.io/terms/

// @host 127.0.0.1:8082
// @BasePath /api/v1
func main() {

	cfg := config.AppConfig()
	log.Println("Starting server...")

	// initialize data sources
	dbs, err := app.InitDS()

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	//Add dependency injection
	router, err := app.InitRouter()

	langLocale := locale.Locale{}
	lang := langLocale.New(cfg.App.Lang)

	appConfig := app.AppConfig{
		Router:  router,
		BaseURL: cfg.App.BaseURL,
		Lang:    lang,
		Dbs:     dbs,
	}

	module.Inject(appConfig)

	if err != nil {
		log.Fatalf("Unable to inject dependencies: %v\n", err)
	}

	router.Run(":" + cfg.Server.Port)
}
