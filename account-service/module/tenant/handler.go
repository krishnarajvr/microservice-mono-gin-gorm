package tenant

import (
	"log"
	"micro/module/tenant/model"
	"micro/module/tenant/service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	logr "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

type Handler struct {
	TenantService service.ITenantService
}

// ListTenants godoc
// @Summary List all existing tenants
// @Description List all existing tenants
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagger.TenantListResponse
// @Router /tenants [get]
func (h *Handler) ListTenants(c *gin.Context) {

	page := common.Paginator(c)
	for k, vals := range c.Request.Header {
		logr.Infof("%s", k)
		for _, v := range vals {
			logr.Infof("\t%s", v)
		}
	}

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

// @Summary Add tenant
// @Produce  json
// @Tags Tenant
// @Param user body model.TenantForm true "Tenant ID"
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
