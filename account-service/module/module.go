package module

import (
	"micro/app"
	"micro/module/tenant"
	"micro/module/user"

	_ "micro/doc"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Inject(appConfig app.AppConfig) {

	user.Inject(appConfig)
	tenant.Inject(appConfig)

	//Init Swagger routes
	appConfig.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
