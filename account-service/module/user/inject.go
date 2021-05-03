package user

import (
	"micro/app"
	"micro/module/user/repo"
	"micro/module/user/service"
	"micro/util/token"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/locale"
)

type HandlerConfig struct {
	R           *gin.Engine
	UserService service.IUserService
	BaseURL     string
	Lang        *locale.Locale
}

//Inject dependencies
func Inject(appConfig app.AppConfig) {

	userRepo := repo.NewUserRepo(appConfig.Dbs.DB)
	tokenRepo := repo.NewTokenRepo(appConfig.Dbs.DB)

	jwtToken := token.New(token.TokenConfig{
		TokenRepo: tokenRepo,
		Cache:     appConfig.Dbs.Cache,
		Config:    appConfig.Cfg,
	})

	userService := service.NewService(service.ServiceConfig{
		UserRepo: userRepo,
		Lang:     appConfig.Lang,
		Token:    jwtToken,
	})

	InitRoutes(HandlerConfig{
		R:           appConfig.Router,
		UserService: userService,
		BaseURL:     appConfig.BaseURL,
		Lang:        appConfig.Lang,
	})
}
