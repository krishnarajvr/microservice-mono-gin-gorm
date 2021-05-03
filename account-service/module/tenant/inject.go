package tenant

import (
	"micro/app"
	"micro/module/tenant/repo"
	"micro/module/tenant/service"
	userRepo "micro/module/user/repo"

	"github.com/krishnarajvr/micro-common/locale"

	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	R             *gin.Engine
	TenantService service.ITenantService
	BaseURL       string
	Lang          *locale.Locale
}

//Inject all dependencies
func Inject(appConfig app.AppConfig) {

	tenantRepo := repo.NewTenantRepo(appConfig.Dbs.DB)
	userRepo := userRepo.NewUserRepo(appConfig.Dbs.DB)

	tenantService := service.NewService(service.ServiceConfig{
		TenantRepo: tenantRepo,
		UserRepo:   userRepo,
		Lang:       appConfig.Lang,
	})

	InitRoutes(HandlerConfig{
		R:             appConfig.Router,
		TenantService: tenantService,
		BaseURL:       appConfig.BaseURL,
		Lang:          appConfig.Lang,
	})
}
