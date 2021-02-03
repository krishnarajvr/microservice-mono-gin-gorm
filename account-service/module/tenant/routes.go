package tenant

import (
	"log"
)

//InitRoutes - initialize routes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		TenantService: c.TenantService,
	}

	// Create service group
	g := c.R.Group(c.BaseURL)

	log.Println("Setting handler")

	g.GET("/tenants/:id", h.GetTenant)
	g.POST("/tenants/:id", h.UpdateTenant)
	g.PATCH("/tenants/:id", h.PatchTenant)
	g.DELETE("/tenants/:id", h.DeleteTenant)
	g.POST("/tenants", h.AddTenant)
	g.GET("/tenants", h.ListTenants)
}
