package product

import (
	"micro/app"
	"micro/module/product/repo"
	"micro/module/product/service"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/locale"
)

type HandlerConfig struct {
	R              *gin.Engine
	ProductService service.IProductService
	BaseURL        string
	Lang           *locale.Locale
}

//Inject dependencies
func Inject(appConfig app.AppConfig) {

	productRepo := repo.NewProductRepo(appConfig.Dbs.DB)

	productService := service.NewService(service.ServiceConfig{
		ProductRepo: productRepo,
		Lang:        appConfig.Lang,
	})

	InitRoutes(HandlerConfig{
		R:              appConfig.Router,
		ProductService: productService,
		BaseURL:        appConfig.BaseURL,
		Lang:           appConfig.Lang,
	})
}
