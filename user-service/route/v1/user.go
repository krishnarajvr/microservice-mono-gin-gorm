package route

import (
	controller "micro/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	posts := r.Group("/users")
	{
		posts.GET("/", controller.UserList)
	}
}
