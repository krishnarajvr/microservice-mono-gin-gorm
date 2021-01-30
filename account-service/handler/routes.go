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
	UserService service.IUserService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R           *gin.Engine
	UserService service.IUserService
	BaseURL     string
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService: c.UserService,
	} // currently has no properties

	// Create an account group
	g := c.R.Group(c.BaseURL)

	log.Println("Setting handler")

	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/token", h.Tokens)
	g.GET("/details", h.Details)
	g.GET("/users", h.ListUsers)

	c.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
