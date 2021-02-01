package user

import (
	"micro/app"
	"micro/module/user/repo"
	"micro/module/user/service"

	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	R           *gin.Engine
	UserService service.IUserService
	BaseURL     string
}

func Inject(appConfig app.AppConfig) {

	userRepo := repo.NewUserRepo(appConfig.Dbs.DB)

	userService := service.NewService(service.ServiceConfig{
		UserRepo: userRepo,
		Lang:     appConfig.Lang,
	})

	InitRoutes(HandlerConfig{
		R:           appConfig.Router,
		UserService: userService,
		BaseURL:     appConfig.BaseURL,
	})
}
