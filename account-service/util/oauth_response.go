package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthError struct {
	Error       string `json:"error" example:"invalid_request"`
	Description string `json:"error_description,omitempty" example:"Request was missing type"`
	URI         string `json:"error_uri,omitempty" example:"{}"`
}

func OauthBadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, OauthError{
		Error:       "invalid_request",
		Description: errorMessage,
	})
}

func OauthUnauthorized(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, OauthError{
		Error:       "invalid_client",
		Description: errorMessage,
	})
}

func OauthUnsupportedGrantType(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, OauthError{
		Error:       "unsupported_grant_type",
		Description: errorMessage,
	})
}

func OauthInternalError(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusInternalServerError, OauthError{
		Error:       "invalid_request",
		Description: errorMessage,
	})
}
