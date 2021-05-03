package user

import (
	"micro/module/user/model"
	"micro/module/user/service"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	"github.com/krishnarajvr/micro-common/locale"
	"github.com/unknwon/com"
)

type Handler struct {
	UserService service.IUserService
	Lang        *locale.Locale
}

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param name query string false "User Name"
// @Param email query string false "User Email"
// @Param pageNumber query int false "Go to a specific page number. Start with 1"
// @Param pageOffset query int false "Go to specific record number. Start with 1. It will override the pageNumber if provided"
// @Param pageSize query string false "Page size for the data"
// @Param pageOrder query string false "Page order. Eg: name desc,createdAt desc"
// @Failure 401 {object} swagdto.Error401
// @Failure 401 {object} swagdto.Error403
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.UserListResponse
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	var filters model.UserFilterList

	log := c.MustGet("log").(*common.MicroLog)
	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	page := common.Paginator(c)
	_ = c.ShouldBindQuery(&filters)

	users, pageResult, err := h.UserService.List(tenantId, page, filters)

	if err != nil {
		log.Message("Failed to list users")
		common.InternalServerError(c, "")
		return
	}

	common.SuccessPageResponse(c, "users", users, pageResult)
}

// @Summary Get a user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.UserResponse
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")

	if !idValid {
		log.Message("Invalid User Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", h.Lang.Get("entity_user")))
		return
	}

	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	user, err := h.UserService.Get(tenantId, id)

	if err != nil {
		log.Message(err)
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", h.Lang.Get("entity_user")))
		return
	}

	if user.ID == 0 {
		log.Message("User with id " + string(id) + " not found")
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", h.Lang.Get("entity_user")))
		return
	}

	common.SuccessResponse(c, "user", user)
}

// @Summary Add user
// @Produce  json
// @Tags User
// @Security ApiKeyAuth
// @Param user body model.UserForm true "User Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.UserResponse
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
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Param user body model.UserForm true "User Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.UserResponse
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

// @Summary Update(Patch) user
// @Produce  json
// @Tags User
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Param user body model.UserForm true "User Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.UserResponse
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
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.UserResponse
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

// @Summary Admin Login API
// @Produce  json
// @Tags Authentication
// @Param tenant body model.LoginForm true "Login Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TokenResponse
// @Router /adminLogin [post]
func (h *Handler) AdminLogin(c *gin.Context) {
	var form model.LoginForm

	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	user, err := h.UserService.Login(&form)

	if err != nil {
		common.BadRequestWithMessage(c, "Invalid User name or password")
		return
	}

	tokenDetail, _ := h.UserService.GetToken(user)

	common.SuccessResponse(c, "token", tokenDetail.ToResponse())
}

// @Summary Token Refresh API
// @Produce  json
// @Tags Authentication
// @Param tenant body model.TokenRefresh true "TokenRefresh Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TokenResponse
// @Router /tokenRefresh [post]
func (h *Handler) TokenRefresh(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var form model.TokenRefresh

	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	tokenDetail, err := h.UserService.RefreshToken(form.RefreshToken)

	if err != nil {
		log.Message(err)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "RefreshToken"))
		return
	}

	common.SuccessResponse(c, "token", tokenDetail.ToResponse())
}

//Authorize  - Private endpoint used to validate token and check the role permission for the API
// Used by gateway and other internal services
func (h *Handler) Authorize(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var resoruceData model.AuthorizeData
	err := c.Bind(&resoruceData)

	if err != nil {
		log.Message(err)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "Resource"))
		return
	}

	token, err := h.UserService.ExtractTokenMetadata(c.Request)

	if err != nil {
		log.Message(err)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "Token"))
		return
	}

	uriData, ok := h.UserService.ParseUri(resoruceData.Uri)

	if !ok {
		log.Message("Invalid Uri : " + resoruceData.Uri)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "ResourceUri"))
		return
	}

	canAccess, pemErr := h.UserService.CheckPermission(token.Roles[0], uriData["module"], uriData["resource"], resoruceData.Method)

	if pemErr != nil {
		log.Message(pemErr)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "RolePermission"))
		return
	}

	if !canAccess {
		log.Message("No Rolepermission : " + resoruceData.Uri)
		common.AccessDenied(c, h.Lang.Get("message_invalid_data", "RolePermission"))
		return
	}

	authRespose := model.AuthRespose{
		TenantId:    strconv.FormatInt(token.TenantId, 10),
		UserId:      strconv.FormatInt(token.UserId, 10),
		ReferenceId: token.ReferenceId,
		Type:        token.Type,
		Subject:     token.Subject,
		Roles:       token.Roles,
	}

	common.SuccessResponse(c, "auth", authRespose)
}
