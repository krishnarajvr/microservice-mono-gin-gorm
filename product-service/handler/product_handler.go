package handler

import (
	"log"

	"micro/common"

	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// ListProducts godoc
// @Summary List all existing products
// @Description List all existing products
// @Tags product
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} common.Error500
// @Success 200 {object} swagdto.ProductResponse
// @Router /products [get]
func (h *Handler) ListProducts(c *gin.Context) {

	products, err := h.ProductService.GetAll()

	if err != nil {
		log.Printf("Failed to sign up product: %v\n", err.Error())
		common.BadRequest(c, "Failed")
		return
	}

	common.Success(c, "products", products)

}

// @Summary Get a single product
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/products/{id} [get]
func (h *Handler) GetProductById(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		common.BadRequest(c, valid.Errors)
		return
	}

	product, err := h.ProductService.GetById(id)
	if err != nil {
		common.BadRequest(c, err)
		return
	}

	common.Success(c, "product", product)
}