package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response defines the api response
type Response struct {
	Status int         `json:"code" example:"200"`
	Data   interface{} `json:"data" example:"{data:{users}}"`
	Error  interface{} `json:"error" example:"{}"`
}

func Success(c *gin.Context, key string, body interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   map[string]interface{}{key: body},
		"error":  "",
	})
}
