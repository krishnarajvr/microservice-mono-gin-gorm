package client

import (
	"micro/app"
	"micro/module/client/repo"
	"micro/module/client/service"
	userRepo "micro/module/user/repo"
	"micro/util/token"

	"github.com/gin-gonic/gin"
	"github.com/krishnarajvr/micro-common/locale"
)

type HandlerConfig struct {
	R             *gin.Engine
	ClientService service.IClientService
	BaseURL       string
	Lang          *locale.Locale
}

//Inject dependencies
func Inject(appConfig app.AppConfig) {
	clientRepo := repo.NewClientRepo(appConfig.Dbs.DB)
	tokenRepo := userRepo.NewTokenRepo(appConfig.Dbs.DB)

	jwtToken := token.New(token.TokenConfig{
		TokenRepo: tokenRepo,
		Cache:     appConfig.Dbs.Cache,
		Config:    appConfig.Cfg,
	})

	clientService := service.NewService(service.ServiceConfig{
		ClientRepo: clientRepo,
		Lang:       appConfig.Lang,
		Token:      jwtToken,
	})

	InitRoutes(HandlerConfig{
		R:             appConfig.Router,
		ClientService: clientService,
		BaseURL:       appConfig.BaseURL,
		Lang:          appConfig.Lang,
	})
}
