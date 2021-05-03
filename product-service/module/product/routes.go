package product

//InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	h := Handler{
		ProductService: c.ProductService,
		Lang:           c.Lang,
	}

	//Set api group
	g := c.R.Group(c.BaseURL)
	g.GET("/products/:id", h.GetProduct)
	g.GET("/products", h.ListProducts)
	g.POST("/products", h.AddProduct)
	g.PATCH("/products/:id", h.PatchProducts)
	g.DELETE("/products/:id", h.DeleteProduct)
}
