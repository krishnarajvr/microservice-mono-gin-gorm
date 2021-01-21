package common

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ValidateForm(c *gin.Context, form interface{}) (status bool, errorData *ErrorData) {
	err := c.Bind(form)
	if err != nil {
		return true, nil
	}

	log := c.MustGet("log").(*logrus.Logger)

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		errorData = &ErrorData{
			Code:    INTERNAL_SERVER_ERROR,
			Message: "Internal Server Error",
		}
		return false, errorData
	}

	if !check {
		log.Info("Validation error")
		log.Info(check)
		log.Info(valid.Errors)
		errorDetails := make([]ErrorDetail, 0)

		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
			errorDetails = append(errorDetails,
				ErrorDetail{
					Code:    err.Key,
					Target:  err.Field,
					Message: err.Message,
				},
			)
		}

		errorData = &ErrorData{
			Code:    BAD_REQUEST,
			Message: "BAD Request",
			Details: errorDetails,
		}

		return false, errorData
	}

	return true, nil
}
