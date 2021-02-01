package user

import (
	"log"
	"micro/module/user/model"
	"micro/module/user/service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	logr "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

type Handler struct {
	UserService service.IUserService
}

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserListResponse
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {

	page := common.Paginator(c)
	for k, vals := range c.Request.Header {
		logr.Infof("%s", k)
		for _, v := range vals {
			logr.Infof("\t%s", v)
		}
	}

	users, pageResult, err := h.UserService.List(page)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		common.BadRequest(c, "Failed")
		return
	}

	common.SuccessPageResponse(c, "users", users, pageResult)

}

// @Summary Get a user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserResponse
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	user, err := h.UserService.Get(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "user", user)
}

// @Summary Add user
// @Produce  json
// @Tags User
// @Param user body model.UserForm true "User ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserResponse
// @Router /users [post]
func (h *Handler) AddUser(c *gin.Context) {
	var form model.UserForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	user, err := h.UserService.Add(&form)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "user", user)
}

// @Summary Update user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Param user body model.UserForm true "User ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserResponse
// @Router /users/{id} [post]
func (h *Handler) UpdateUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.UserForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	user, err := h.UserService.Update(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "user", user)
}

// @Summary Patch user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Param user body model.UserForm true "User ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserResponse
// @Router /users/{id} [patch]
func (h *Handler) PatchUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.UserPatchForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	user, err := h.UserService.Patch(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "user", user)
}

// @Summary Delete a user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.UserResponse
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	user, err := h.UserService.Delete(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "user", user)
}
