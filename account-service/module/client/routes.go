package client

//InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		ClientService: c.ClientService,
		Lang:          c.Lang,
	}

	//Set api group
	g := c.R.Group(c.BaseURL)
	g.GET("/clientCredentials/:id", h.GetClientCredential)
	g.GET("/clientCredentials", h.ListClientCredentials)
	g.POST("/clientCredentials", h.AddClientCredential)
	g.PATCH("/clientCredentials/:id", h.PatchClientCredential)
	g.POST("/clientLogin", h.ClientLogin)
	g.POST("/oauth/token", h.OauthToken)
}
