package app

import (
	"log"

	"micro/app/locale"
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
	userRepo := repo.NewUserRepo(d.DB)
	langLocale := locale.Locale{}
	lang := langLocale.New(appConf.App.Lang)

	userService := service.NewUserService(&service.ServiceConfig{
		UserRepo: userRepo,
		Lang:     lang,
	})

	// initialize gin.Engine
	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R:           router,
		UserService: userService,
		BaseURL:     appConf.App.BaseURL,
	})

	return router, nil
}
