package common

import (
	"micro/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginator(c *gin.Context) model.Pagination {
	var page model.Pagination

	// Define and get sorting field
	page.Sort = c.DefaultQuery("sort", "ID")

	// Define and get sorting order field
	page.Order = c.DefaultQuery("order", "DESC")

	// Define and get offset for pagination
	page.Offset = c.DefaultQuery("offset", "0")

	// Define and get limit for pagination
	page.Limit = c.DefaultQuery("limit", "25")

	// Get search keyword for Search Scope
	page.Search = c.DefaultQuery("Search", "")

	return page
}

func Paginate(page model.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		limitNumber, _ := strconv.Atoi(page.Limit)
		offsetNumber, _ := strconv.Atoi(page.Offset)

		offset := (offsetNumber - 1) * limitNumber
		return db.Offset(offset).Limit(limitNumber).Order(page.Sort + " " + page.Order)
	}
}
