package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationError(err error, c *gin.Context) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(map[string]string)
		for _, fieldError := range ve {
			fieldName := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				out[fieldName] = "This field is required"
			default:
				out[fieldName] = fieldName + " is not valid"
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error :" + err.Error()})
	}
}
