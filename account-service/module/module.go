package module

import (
	"micro/app"
	"micro/module/tenant"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Inject(appConfig app.AppConfig) {

	tenant.Inject(appConfig)

	//Init Swagger routes
	appConfig.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
