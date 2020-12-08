package controller

import (
	"micro/service"

	"github.com/gin-gonic/gin"
)

// UserList godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} common.Error500
// @Success 200 {object} swagdto.UserResponse
// @Router /users [get]
func UserList(c *gin.Context) {

	users, err := service.ListUsers(c)

	if err != nil {
		c.JSON(500, "error")
		return
	}

	c.JSON(200, users)
}
