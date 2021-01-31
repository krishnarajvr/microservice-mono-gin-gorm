package tenant

import (
	"micro/app"
	"micro/module/tenant/repo"
	"micro/module/tenant/service"

	"github.com/gin-gonic/gin"
)

type HandlerConfig struct {
	R             *gin.Engine
	TenantService service.ITenantService
	BaseURL       string
}

func Inject(appConfig app.AppConfig) {

	tenantRepo := repo.NewTenantRepo(appConfig.Dbs.DB)

	tenantService := service.NewService(service.ServiceConfig{
		TenantRepo: tenantRepo,
		Lang:       appConfig.Lang,
	})

	InitRoutes(HandlerConfig{
		R:             appConfig.Router,
		TenantService: tenantService,
		BaseURL:       appConfig.BaseURL,
	})
}
