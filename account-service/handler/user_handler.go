package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	common "github.com/krishnarajvr/micro-common"
)

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} common.Error500
// @Success 200 {object} swagdto.UserResponse
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {

	users, err := h.UserService.GetAll()

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		common.BadRequest(c, "No users")
		return
	}
	common.SuccessResponse(c, "users", users)

}
