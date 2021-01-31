package tenant

import (
	"log"
)

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func InitRoutes(c HandlerConfig) {
	// Create a handler (which will later have injected services)
	h := Handler{
		TenantService: c.TenantService,
	} // currently has no properties

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
