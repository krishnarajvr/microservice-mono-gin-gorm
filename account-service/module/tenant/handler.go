package tenant

import (
	"encoding/json"
	"log"
	"micro/module/tenant/model"
	"micro/module/tenant/service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	"github.com/krishnarajvr/micro-common/locale"
	"github.com/unknwon/com"
)

type Handler struct {
	TenantService service.ITenantService
	Lang          *locale.Locale
}

// ListTenants godoc
// @Summary List all existing tenants
// @Description List all existing tenants
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantListResponse
// @Router /tenants [get]
func (h *Handler) ListTenants(c *gin.Context) {
	page := common.Paginator(c)

	tenants, err := h.TenantService.List(page)

	if err != nil {
		log.Printf("Failed to sign up tenant: %v\n", err.Error())
		common.BadRequest(c, "Failed")
		return
	}

	common.SuccessResponse(c, "tenants", tenants)
}

// @Summary Get a tenant
// @Produce  json
// @Tags Tenant
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenants/{id} [get]
func (h *Handler) GetTenant(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	tenant, err := h.TenantService.Get(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

// @Summary Register tenant
// @Produce  json
// @Tags Tenant
// @Param user body model.TenantRegisterForm true "Tenant data"
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenantRegister [post]
func (h *Handler) RegisterTenant(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var form model.TenantRegisterForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	tenant, err := h.TenantService.Register(&form)
	if err != nil {
		log.Message(err)
		errorCode := common.CheckDbError(err)

		if errorCode == common.ALREADY_EXISTS {
			common.BadRequestWithMessage(c, h.Lang.Get("message_already_exists", "Tenant"))
			return
		}

		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))

		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

// @Summary Add tenant
// @Produce  json
// @Tags Tenant
// @Security ApiKeyAuth
// @Param user body model.TenantForm true "Tenant Data"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenants [post]
func (h *Handler) AddTenant(c *gin.Context) {
	var form model.TenantForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	tenant, err := h.TenantService.Add(&form)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

// @Summary Update tenant
// @Produce  json
// @Tags Tenant
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Param user body model.TenantForm true "Tenant ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenants/{id} [post]
func (h *Handler) UpdateTenant(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.TenantForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	tenant, err := h.TenantService.Update(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

// @Summary Patch tenant
// @Produce  json
// @Tags Tenant
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Param user body model.TenantForm true "Tenant ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenants/{id} [patch]
func (h *Handler) PatchTenant(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.TenantPatchForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	tenant, err := h.TenantService.Patch(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

// @Summary Delete a tenant
// @Produce  json
// @Tags Tenant
// @Param id path int true "ID"
// @Security ApiKeyAuth
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantResponse
// @Router /tenants/{id} [delete]
func (h *Handler) DeleteTenant(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	tenant, err := h.TenantService.Delete(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "tenant", tenant)
}

//GetJwks - Return Keys for JWT
func (h *Handler) GetJwks(c *gin.Context) {
	//Todo - Get values from env or external file
	keys := `{"keys": [
		{
			"kty": "oct",
			"alg": "A128KW",
			"k": "GawgguFyGrWKav7AX4VKUg",
			"kid": "sim1"
		},
		{
			"kty": "oct",
			"k": "AyM1SysPpbyDfgZld3umj1qzKObwVMkoqQ-EstJQLr_T-1qS0gZH75aKtMN3Yj0iPS4hcgUuTwjAzZr1Z9CAow",
			"kid": "sim2",
			"alg": "HS256"
		}
	]}`

	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(keys), &result)

	c.JSON(200, result)
}
