package repositories

import (
	"data-parameter/config"
	"data-parameter/models"
)

func ResponseMessageFindByCodeAndSorce(code string, source string) (*models.ResponseMessage, error) {
	var responseMessage models.ResponseMessage
	result := config.DB.First(&responseMessage, "code = $1 AND source = $2", code, source)
	return &responseMessage, result.Error
}
