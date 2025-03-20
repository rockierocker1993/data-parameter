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

func GetAllResponseMessages() ([]models.ResponseMessage, error) {
	var responseMessages []models.ResponseMessage
	result := config.DB.Find(&responseMessages)
	return responseMessages, result.Error
}

func CreateResponseMessage(responseMessage *models.ResponseMessage) error {
	result := config.DB.Create(responseMessage)
	return result.Error
}

func GetResponseMessageByID(id uint) (*models.ResponseMessage, error) {
	var responseMessage models.ResponseMessage
	result := config.DB.First(&responseMessage, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &responseMessage, nil
}

func GetResponseMessageByCode(code string) (*models.ResponseMessage, error) {
	var responseMessage models.ResponseMessage
	result := config.DB.First(&responseMessage, "code = ?", code)
	if result.Error != nil {
		return nil, result.Error
	}
	return &responseMessage, nil
}

func IsExistResponseMessageByCodeAndIdNot(code string, id int) error {
	var responseMessage models.ResponseMessage
	result := config.DB.First(&responseMessage, "code = $1 AND ID !=$2", code, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateResponseMessage(responseMessage *models.ResponseMessage) error {
	result := config.DB.Save(responseMessage)
	return result.Error
}

func DeleteResponseMessage(id uint) error {
	result := config.DB.Delete(&models.ResponseMessage{}, id)
	return result.Error
}
