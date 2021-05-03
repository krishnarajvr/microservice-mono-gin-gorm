package user

//InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		UserService: c.UserService,
		Lang:        c.Lang,
	}
	//Set api group
	g := c.R.Group(c.BaseURL)

	g.GET("/users/:id", h.GetUser)
	g.POST("/users/:id", h.UpdateUser)
	g.PATCH("/users/:id", h.PatchUser)
	g.DELETE("/users/:id", h.DeleteUser)
	g.POST("/users", h.AddUser)
	g.GET("/users", h.ListUsers)
	g.POST("/adminLogin", h.AdminLogin)
	g.POST("/tokenRefresh", h.TokenRefresh)
	g.POST("/authorize", h.Authorize)
}
