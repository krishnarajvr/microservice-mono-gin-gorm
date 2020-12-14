package handler

import (
	"log"
	"net/http"

	_ "micro/doc"
	"micro/service"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Handler struct holds required services for handler to function
type Handler struct {
	ProductService service.IProductService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R              *gin.Engine
	ProductService service.IProductService
	BaseURL        string
}

func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		ProductService: c.ProductService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(c.BaseURL)

	log.Println("Setting handler")

	g.GET("/products/:id", h.GetProductById)
	g.GET("/products", h.ListProducts)

	c.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
