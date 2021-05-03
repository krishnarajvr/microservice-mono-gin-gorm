package tenant

//InitRoutes - initialize routes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		TenantService: c.TenantService,
		Lang:          c.Lang,
	}

	// Create service group
	g := c.R.Group(c.BaseURL)

	g.GET("/tenants/:id", h.GetTenant)
	g.POST("/tenants/:id", h.UpdateTenant)
	g.PATCH("/tenants/:id", h.PatchTenant)
	g.DELETE("/tenants/:id", h.DeleteTenant)
	g.POST("/tenants", h.AddTenant)
	g.POST("/tenantRegister", h.RegisterTenant)
	g.GET("/tenants", h.ListTenants)
	g.GET("/jwks", h.GetJwks)
}
