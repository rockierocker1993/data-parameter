package util

import (
	"data-parameter/constant"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

func ValidateRequest[T any](c *gin.Context, payload *T) bool {
	if err := c.ShouldBindJSON(payload); err != nil {
		slog.Debug("Invalid request payload", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, ConstructResponse(c, constant.INVALID_REQUEST, constant.Source, nil))
		return false
	}

	// Validasi field berdasarkan tag struct
	if err := validate.Struct(payload); err != nil {
		slog.Debug("Validation failed", slog.Any("error", err))
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}

		c.JSON(http.StatusBadRequest, ConstructResponse(c, constant.VALIDATION_FAILED, constant.Source, nil))
		return false
	}
	return true
}

func ValidateRequestSingleField[T any](c *gin.Context, payload *T, prefix string) bool {
	if err := c.ShouldBindJSON(payload); err != nil {
		slog.Debug("Invalid request payload", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, ConstructResponse(c, constant.INVALID_REQUEST, constant.Source, nil))
		return false
	}

	// Validasi field berdasarkan tag struct
	if err := validate.Struct(payload); err != nil {
		slog.Debug("Validation failed", slog.Any("error", err))
		for _, err := range err.(validator.ValidationErrors) {
			errorCode := strings.ToUpper(prefix + "_" + err.Field() + "_" + err.Tag())
			c.JSON(http.StatusBadRequest, ConstructResponse(c, errorCode, constant.Source, nil))
			return false
		}
	}
	return true
}
