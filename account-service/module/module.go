package module

import (
	"micro/app"
	"micro/module/client"
	"micro/module/tenant"
	"micro/module/user"
	"os"

	docs "micro/doc"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Micro Service API Document
// @version 1.0
// @description List of APIs for Micro Service
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8082
// @BasePath /api/v1
func Inject(appConfig app.AppConfig) {
	user.Inject(appConfig)
	tenant.Inject(appConfig)
	client.Inject(appConfig)

	//Swagger Doc details
	url := os.Getenv("API_GATEWAY_URL")
	prefix := os.Getenv("API_GATEWAY_PREFIX")

	if len(url) == 0 {
		url = "localhost:" + appConfig.Cfg.Server.Port
	}

	if len(prefix) == 0 {
		prefix = appConfig.Cfg.App.BaseURL
	}

	docs.SwaggerInfo.Title = "Account Service API Document"
	docs.SwaggerInfo.Description = "List of APIs for Account Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = url
	docs.SwaggerInfo.BasePath = prefix
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	//Init Swagger routes
	appConfig.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
