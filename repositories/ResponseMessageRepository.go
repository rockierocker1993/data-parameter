package repositories

import (
	"data-parameter/config"
	"data-parameter/models"
)

func ResponseMessageFindByCodeAndSorce(code string, source string) (*models.ResponseMessage, error) {
	var responseMessage models.ResponseMessage
	result := config.DB.First(&responseMessage, "code = ?", code, "source = ?", source)
	return &responseMessage, result.Error
}
