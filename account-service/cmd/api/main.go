package main

import (
	"log"

	"micro/app"
	"micro/config"
	"micro/module"
)

func main() {
	cfg := config.AppConfig()

	log.Println("Starting server...")

	// initialize data sources
	dbs, err := app.InitDS(cfg)

	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	//Close the database connection when stopped
	defer app.Close(dbs)

	//Add dependency injection

	//Public routes that don't have tenant checking
	excludeList := map[string]interface{}{
		"/api/v1/token":           true,
		"/health":                 true,
		"/api/v1/tenantRegister":  true,
		"/api/v1/adminLogin":      true,
		"/api/v1/clientLogin":     true,
		"/api/v1/rolePermissions": true,
		"/api/v1/authorize":       true,
		"/api/v1/tokenRefresh":    true,
		"/api/v1/oauth/token":     true,
	}
	router, err := app.InitRouter(cfg, excludeList)

	if err != nil {
		log.Fatalf("Unable to initialize routes: %v\n", err)
	}

	lang, err := app.InitLocale(cfg)

	if err != nil {
		log.Fatalf("Unable to initialize language locale: %v\n", err)
	}

	appConfig := app.AppConfig{
		Router:  router,
		BaseURL: cfg.App.BaseURL,
		Lang:    lang,
		Dbs:     dbs,
		Cfg:     cfg,
	}

	module.Inject(appConfig)

	if err != nil {
		log.Fatalf("Unable to inject dependencies: %v\n", err)
	}

	router.Run(":" + cfg.Server.Port)
}
