package handler

import (
	"log"

	"micro/common"

	"github.com/gin-gonic/gin"
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
		c.JSON(common.Status(err), gin.H{
			"error": err,
		})
		return
	}

	common.Success(c, "products", products)

}
