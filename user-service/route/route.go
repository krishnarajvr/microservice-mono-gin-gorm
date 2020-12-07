package route

import (
	_ "micro/doc"
	"micro/route/v1"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Apply(r *gin.Engine) {
	api := r.Group("/v1")
	{
		route.UserRoute(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
