package handler

import (
	"log"
	"micro/model"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"

	common "github.com/krishnarajvr/micro-common"
	logr "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// ListProducts godoc
// @Summary List all existing products
// @Description List all existing products
// @Tags product
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductListResponse
// @Router /products [get]
func (h *Handler) ListProducts(c *gin.Context) {

	page := common.Paginator(c)
	for k, vals := range c.Request.Header {
		logr.Infof("%s", k)
		for _, v := range vals {
			logr.Infof("\t%s", v)
		}
	}

	products, err := h.ProductService.List(page)

	if err != nil {
		log.Printf("Failed to sign up product: %v\n", err.Error())
		common.BadRequest(c, "Failed")
		return
	}

	common.SuccessResponse(c, "products", products)

}

// @Summary Get a product
// @Produce  json
// @Param id path int true "ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductResponse
// @Router /products/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	product, err := h.ProductService.Get(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Add product
// @Produce  json
// @Param user body model.ProductForm true "Product ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductResponse
// @Router /products [post]
func (h *Handler) AddProduct(c *gin.Context) {
	var form model.ProductForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	product, err := h.ProductService.Add(&form)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Update product
// @Produce  json
// @Param id path int true "ID"
// @Param user body model.ProductForm true "Product ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductResponse
// @Router /products/{id} [post]
func (h *Handler) UpdateProduct(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.ProductForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	product, err := h.ProductService.Update(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Patch product
// @Produce  json
// @Param id path int true "ID"
// @Param user body model.ProductForm true "Product ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductResponse
// @Router /products/{id} [patch]
func (h *Handler) PatchProduct(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	var form model.ProductPatchForm

	ok, errorData := common.ValidateForm(c, &form)
	if !ok {
		common.ErrorResponse(c, errorData)
		return
	}

	product, err := h.ProductService.Patch(&form, id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "product", product)
}

// @Summary Delete a product
// @Produce  json
// @Param id path int true "ID"
// @Failure 404 {object} swagdto.Error404
// @Success 200 {object} swagdto.ProductResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	product, err := h.ProductService.Delete(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.SuccessResponse(c, "product", product)
}
