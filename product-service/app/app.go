package app

import (
	"log"

	"micro/app/locale"
	"micro/app/middleware"
	"micro/config"
	"micro/handler"
	"micro/service"

	"micro/repo"

	"github.com/gin-gonic/gin"
)

func Inject(d *dbs) (*gin.Engine, error) {
	log.Println("Injecting common classes")

	//Initialize all common classes
	appConf := config.AppConfig()
	productRepo := repo.NewProductRepo(d.DB)
	langLocale := locale.Locale{}
	lang := langLocale.New(appConf.App.Lang)

	productService := service.NewProductService(&service.ServiceConfig{
		ProductRepo: productRepo,
		Lang:        lang,
	})

	// initialize gin.Engine
	router := gin.Default()
	router.Use(middleware.LoggerToFile())
	router.Use(gin.Recovery())

	handler.NewHandler(&handler.Config{
		R:              router,
		ProductService: productService,
		BaseURL:        appConf.App.BaseURL,
	})

	return router, nil
}
