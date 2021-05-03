package product

import (
	"encoding/json"
	"micro/module/product/model"
	"micro/module/product/service"
	"strconv"

	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	"github.com/krishnarajvr/micro-common/locale"
)

type Handler struct {
	ProductService service.IProductService
	Lang           *locale.Locale
}

// ListProducts godoc
// @Summary List all existing products
// @Description List all existing products
// @Tags Product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param code query string false "filter the code based value (ex: code=1000)"
// @Param name query string false "filter the name based value (ex: name=Ja)"
// @Param pageNumber query int false "Go to a specific page number. Start with 1"
// @Param pageOffset query int false "Go to specific record number. Start with 1. It will override the pageNumber if provided"
// @Param pageSize query string false "Page size for the data"
// @Param pageOrder query string false "Page order. Eg: name desc,createdAt desc"
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ProductListResponse
// @Router /products [get]
func (h *Handler) ListProducts(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	var filters model.ProductFilterList
	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))

	page := common.Paginator(c)

	filters.Code = c.DefaultQuery("code", "")
	filters.Name = c.DefaultQuery("name", "")

	products, pageResult, err := h.ProductService.List(tenantId, page, filters)

	if err != nil {
		log.Message("Failed to list content" + err.Error())
		common.InternalServerError(c, "")
		return
	}
	common.SuccessPageResponse(c, "product", products, pageResult)
}

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product
// @Tags Product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param product body swagger.ProductForm true "Product Data"
// @Failure 400 {object} swagdto.Error400
// @Failure 403 {object} swagdto.Error403
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ProductAddResponse
// @Router /products [post]
func (h *Handler) AddProduct(c *gin.Context) {
	var form model.ProductForm
	log := c.MustGet("log").(*common.MicroLog)
	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	ok, errorData := common.ValidateForm(c, &form)

	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}
	if form.Meta != nil {

		addmeta, err := json.Marshal(form.Meta)

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		metaschemaDef, metaformData, err := h.ProductService.CheckProductSchema(tenantId, addmeta, "validateMetaData")

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		validationConfig := map[string]string{
			"trimRequired": "true",
		}

		formMetaTrimmed, errorData := common.ValidateCustomSchema(&metaschemaDef, metaformData, validationConfig)

		if errorData != nil {
			log.Message(errorData)
			common.ErrorResponse(c, errorData)
			return

		}

		if validationConfig["trimRequired"] != "false" {
			metaJson, err := h.ProductService.ConvertMapToJson(formMetaTrimmed)

			if err != nil {
				log.Message(err)
				common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
				return
			}

			form.Meta = metaJson
		}

	}
	if form.Meta != nil {

		addaddress, err := json.Marshal(form.Meta)

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		addressschemaDef, addressformData, err := h.ProductService.CheckProductSchema(tenantId, addaddress, "validateMetaData")

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		validationConfig := map[string]string{
			"trimRequired": "true",
		}

		formMetaTrimmed, errorData := common.ValidateCustomSchema(&addressschemaDef, addressformData, validationConfig)

		if errorData != nil {
			log.Message(errorData)
			common.ErrorResponse(c, errorData)
			return

		}

		if validationConfig["trimRequired"] != "false" {
			addressJson, err := h.ProductService.ConvertMapToJson(formMetaTrimmed)

			if err != nil {
				log.Message(err)
				common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
				return
			}

			form.Meta = addressJson
		}
	}

	product, err := h.ProductService.Add(tenantId, &form)

	if err != nil {
		log.Message(err)
		errorCode := common.CheckDbError(err)

		if errorCode == common.ALREADY_EXISTS {
			common.BadRequestWithMessage(c, h.Lang.Get("message_already_exists", "Product"))
			return
		}

		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))

		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Get a product
// @Description Get a specific product by id
// @Produce  json
// @Tags Product
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Failure 403 {object} swagdto.Error403
// @Failure 404 {object} swagdto.Error404
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ProductResponse
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")

	if !idValid {
		log.Message("Invalid Product Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "Product"))
		return
	}

	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))
	product, err := h.ProductService.Get(tenantId, id)

	if err != nil {
		log.Message(err)
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "Product"))
		return
	}

	if product.ID == 0 {
		log.Message("Product with id " + strconv.Itoa(id) + " not found")
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "Product"))
		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Update a product
// @Description Update a specific product
// @Produce json
// @Tags Product
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Param product body swagger.ProductPatchForm true "Product Data"
// @Failure 400 {object} swagdto.Error400
// @Failure 403 {object} swagdto.Error403
// @Failure 404 {object} swagdto.Error404
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ProductUpdateResponse
// @Router /products/{id} [patch]
func (h *Handler) PatchProducts(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")
	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))

	if !idValid {
		log.Message("Invalid Product Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "Product"))
		return
	}

	var patchform model.ProductPatchForm
	ok, errorData := common.ValidateForm(c, &patchform)

	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	productRes, _ := h.ProductService.Get(tenantId, id)

	if productRes == nil {
		log.Message("Invalid Product Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "Product"))
		return
	}

	if patchform.Meta != nil {

		patchmeta, err := json.Marshal(patchform.Meta)

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		metaschemaDef, metaformData, err := h.ProductService.CheckProductSchema(tenantId, patchmeta, "validateMetaData")

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		validationConfig := map[string]string{
			"trimRequired": "true",
		}

		formMetaTrimmed, errorData := common.ValidateCustomSchema(&metaschemaDef, metaformData, validationConfig)

		if errorData != nil {
			log.Message(errorData)
			common.ErrorResponse(c, errorData)
			return

		}

		if validationConfig["trimRequired"] != "false" {
			metaJson, err := h.ProductService.ConvertMapToJson(formMetaTrimmed)

			if err != nil {
				log.Message(err)
				common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
				return
			}

			patchform.Meta = metaJson
		}

	}

	if patchform.Meta != nil {

		patchaddress, err := json.Marshal(patchform.Meta)

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		addressschemaDef, addressformData, err := h.ProductService.CheckProductSchema(tenantId, patchaddress, "validateMetaData")

		if err != nil {
			log.Message("convert jsonschema fail: " + err.Error())
			common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
			return
		}

		validationConfig := map[string]string{
			"trimRequired": "true",
		}

		formMetaTrimmed, errorData := common.ValidateCustomSchema(&addressschemaDef, addressformData, validationConfig)

		if errorData != nil {
			log.Message(errorData)
			common.ErrorResponse(c, errorData)
			return
		}

		if validationConfig["trimRequired"] != "false" {
			addressJson, err := h.ProductService.ConvertMapToJson(formMetaTrimmed)

			if err != nil {
				log.Message(err)
				common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
				return
			}

			patchform.Meta = addressJson
		}

	}

	_, errPatch := h.ProductService.Patch(&patchform, id)

	//Todo - Handle different db conditions with proper messages
	if errPatch != nil {
		log.Message("Product with id " + strconv.Itoa(id) + " not found")
		common.ResourceNotFound(c, h.Lang.Get("message_not_found", "product"))
		return

	}

	product, err := h.ProductService.Get(tenantId, id)

	if err != nil {

		log.Message(err)

		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))

		return

	}

	product.ID = uint(id)
	common.SuccessResponse(c, "product", product)
}

// @Summary Delete a product
// @Description Delete a specific product by id
// @Produce  json
// @Tags Product
// @Param id path int true "Product ID"
// @Security ApiKeyAuth
// @Failure 403 {object} swagdto.Error403
// @Failure 404 {object} swagdto.Error404
// @Failure 500 {object} swagdto.Error500
// @Success 200 {object} swagger.ProductResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	log := c.MustGet("log").(*common.MicroLog)
	id, idValid := common.ValidateID(c, "id")

	if !idValid {
		log.Message("Invalid Product Id")
		common.BadRequestWithMessage(c, h.Lang.Get("message_invalid_id", "Product"))
		return
	}

	tenantId, _ := strconv.Atoi(c.MustGet("tenantId").(string))

	product, err := h.ProductService.Delete(tenantId, id)

	if err != nil {
		log.Message("service returned error: " + err.Error())
		common.InternalServerError(c, h.Lang.Get("message_internal_error", ""))
		return
	}

	common.SuccessResponse(c, "product", product)
}
