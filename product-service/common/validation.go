package common

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ValidateForm(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, INVALID_PARAMS
	}

	log := c.MustGet("log").(*logrus.Logger)

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	
	if err != nil {
		return http.StatusInternalServerError, ERROR
	}

	if !check {
		log.Info("Validation error")
		log.Info(check)
		log.Info(valid.Errors)
		return http.StatusBadRequest, INVALID_PARAMS
	}

	return http.StatusOK, SUCCESS
}
