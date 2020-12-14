package handler

import (
	"log"
	"net/http"

	"micro/service"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	ProductService service.IProductService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R           *gin.Engine
	ProductService service.IProductService
	BaseURL     string
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		ProductService: c.ProductService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(c.BaseURL)

	log.Println("Setting handler")

	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.GET("/details", h.Details)
	g.GET("/users", h.ListProducts)
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
