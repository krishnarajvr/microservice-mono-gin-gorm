package user

import (
	"log"
)

//InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		UserService: c.UserService,
	}

	g := c.R.Group(c.BaseURL)

	log.Println("Setting handler")

	g.GET("/users/:id", h.GetUser)
	g.POST("/users/:id", h.UpdateUser)
	g.PATCH("/users/:id", h.PatchUser)
	g.DELETE("/users/:id", h.DeleteUser)
	g.POST("/users", h.AddUser)
	g.GET("/users", h.ListUsers)
}
