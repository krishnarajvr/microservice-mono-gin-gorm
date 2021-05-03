package client

import (
	"micro/module/client/model"
	"micro/module/client/service"
	"micro/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	"github.com/krishnarajvr/micro-common/locale"
)

type Handler struct {
	ClientService service.IClientService
	Lang          *locale.Locale
}

// ListClientCredentials godoc
// @Summary List all existing client credentials
// @Description List all existing client credentials
// @Tags Client
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param pageNumber query int false "Go to a specific page number. Start with 1"
// @Param pageOffset query int false "Go to specific record number. Start with 1. It will override the pageNumber if provided"
// @Param pageSize query string false "Page size for the data"
// @Param pageOrder query string false "Page order. Eg: name desc,createdAt desc"
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ClientListCredentialsResponse
// @Router /clientCredentials [get]
func (h *Handler) ListClientCredentials(c *gin.Context) {
	microLog := c.MustGet("log").(*common.MicroLog)
	page := common.Paginator(c)

	clients, pageResult, err := h.ClientService.ListCredentials(page)

	if err != nil {
		microLog.Message(err)
		common.BadRequestWithMessage(c, h.Lang.Get("message_internal_error", ""))
		return
	}

	common.SuccessPageResponse(c, "clientCredentials", clients, pageResult)
}

// AddClientCredential godoc
// @Summary Add a new client
// @Description Add a new client credential
// @Tags Client
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param client body swagger.ClientCredentialForm true "ClientCredential Data"
// @Failure 400 {object} swagdto.Error400
// @Failure 403 {object} swagdto.Error403
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ClientListCredentialsResponse
// @Router /clientCredentials [post]
func (h *Handler) AddClientCredential(c *gin.Context) {
	var form model.ClientCredentialForm
	log := c.MustGet("log").(*common.MicroLog)
	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	// As of now its accepting the "vendor", future it will enhance further.
	if strings.ToLower(form.Type) != "vendor" {
		log.Message("Invalid type" + form.Type)
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_data", "type"))
		return
	}

	client, err := h.ClientService.AddCredential(&form, tenantId)

	if err != nil {
		log.Message(err)
		errorCode := common.CheckDbError(err)

		if errorCode == common.ALREADY_EXISTS {
			common.BadRequestWithMessage(c, h.Lang.Get("message_already_exists", "ClientCredential"))
			return
		}

		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))

		return
	}

	common.SuccessResponse(c, "clientCredential", client)
}

// @Summary Get client credential
// @Description Get a specific client credential by id
// @Produce  json
// @Tags Client
// @Param id path int true "ClientCredential ID"
// @Security ApiKeyAuth
// @Failure 403 {object} swagdto.Error403
// @Failure 404 {object} swagdto.Error404
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ClientCredentialResponse
// @Router /clientCredentials/{id} [get]
func (h *Handler) GetClientCredential(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")

	if !idValid {
		log.Message("Invalid ClientCredential Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "ClientCredential"))
		return
	}

	//Todo: Handle tenant Id in all the requests
	//tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	client, err := h.ClientService.GetCredential(id)

	if err != nil {
		log.Message(err)
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "ClientCredential"))
		return
	}

	if client.ID == 0 {
		log.Message("ClientCredential with id " + strconv.Itoa(id) + " not found")
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "ClientCredential"))
		return
	}

	common.SuccessResponse(c, "clientCredential", client)
}

// @Summary Update client credential
// @Description Update a specific client
// @Produce json
// @Tags Client
// @Param id path int true "ClientCredential ID"
// @Security ApiKeyAuth
// @Param client body swagger.ClientCredentialForm true "ClientCredential Data"
// @Failure 400 {object} swagdto.Error400
// @Failure 403 {object} swagdto.Error403
// @Failure 404 {object} swagdto.Error404
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ClientCredentialResponse
// @Router /clientCredentials/{id} [patch]
func (h *Handler) PatchClientCredential(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")

	if !idValid {
		log.Message("Invalid ClientCredential Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "ClientCredential"))
		return
	}

	var form model.ClientCredentialPatchForm
	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	client, err := h.ClientService.PatchCredential(&form, id)

	//Todo - Handle different db conditions with proper messages
	if err != nil {
		log.Message("ClientCredential with id " + strconv.Itoa(id) + " not found")
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "client"))
		return

	}

	client.ID = uint(id)
	common.SuccessResponse(c, "clientCredential", client)
}

func (h *Handler) OauthToken(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var form model.ClientLoginForm

	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		util.OauthBadRequest(c, errorData.Details[0].Message)
		return
	}

	if form.GrantType != "client_credentials" && form.GrantType != "refresh_token" {
		log.Message("Invalid Type:" + form.GrantType)
		util.OauthBadRequest(c, h.Lang.Get("message_invalid_data", "grant_type"))

		return
	}

	clientCredential, err := h.ClientService.CheckCredentials(&form)

	if err != nil {
		log.Message("Invalid Client Credentials: " + err.Error())
		util.OauthUnauthorized(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	if len(clientCredential.Name) == 0 {
		log.Message("Empty client credentials")
		util.OauthUnauthorized(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	roles, errC := h.ClientService.GetClientCredentialRoles(int(clientCredential.ID))

	if errC != nil {
		log.Message("No role mapping found for clients")
		util.OauthUnauthorized(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	if form.GrantType == "client_credentials" {
		tokenDetail, errT := h.ClientService.GetClientToken(clientCredential, roles)

		if errT != nil {
			log.Message(errT)
			util.OauthInternalError(c, h.Lang.Get("message_internal_error", ""))

			return
		}

		c.JSON(200, tokenDetail.OAuthResponse())

		return
	}

	if form.GrantType == "refresh_token" {
		tokenDetail, errT := h.ClientService.RefreshToken(form.RefreshToken)

		if errT != nil {
			log.Message(errT)
			util.OauthInternalError(c, h.Lang.Get("message_internal_error", ""))

			return
		}

		c.JSON(200, tokenDetail.OAuthResponse())

		return
	}

	util.OauthBadRequest(c, h.Lang.Get("message_invalid_data", "request"))
}

func (h *Handler) ClientLogin(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var form model.ClientLoginForm

	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		common.ErrorResponse(c, errorData)

		return
	}

	if strings.ToLower(form.GrantType) != "client_credentials" {
		log.Message("Invalid Type:" + form.GrantType)
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_data", "grant_type"))

		return
	}

	clientCredential, err := h.ClientService.CheckCredentials(&form)

	if err != nil {
		log.Message("Invalid Client Credentials: " + err.Error())
		common.AccessDenied(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	if len(clientCredential.Name) == 0 {
		log.Message("Empty client credentials")
		common.AccessDenied(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	roles, errC := h.ClientService.GetClientCredentialRoles(int(clientCredential.ID))

	if errC != nil {
		log.Message("No role mapping found for clients")
		common.AccessDenied(c, h.Lang.Get("message_invalid_client_credentials"))

		return
	}

	tokenDetail, errT := h.ClientService.GetClientToken(clientCredential, roles)

	if errT != nil {
		log.Message(errT)
		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))

		return
	}

	common.SuccessResponse(c, "token", tokenDetail.ToResponse())
}
