package module

import (
	"micro/app"
	product "micro/module/product"
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

// @host localhost:8083
// @BasePath /api/v1
func Inject(appConfig app.AppConfig) {
	product.Inject(appConfig)

	//Swagger Doc details
	url := os.Getenv("API_GATEWAY_URL")
	prefix := os.Getenv("API_GATEWAY_PREFIX")

	if len(url) == 0 {
		url = "localhost:" + appConfig.Cfg.Server.Port
	}

	if len(prefix) == 0 {
		prefix = appConfig.Cfg.App.BaseURL
	}

	docs.SwaggerInfo.Title = "Product Service API Document"
	docs.SwaggerInfo.Description = "List of APIs for Product Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = url
	docs.SwaggerInfo.BasePath = prefix
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	//Init Swagger routes
	appConfig.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
