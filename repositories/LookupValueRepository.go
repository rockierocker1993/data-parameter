package repositories

import (
	"data-parameter/config"
	"data-parameter/models"
)

func GetAllLookupValues() ([]models.LookupValue, error) {
	var lookupValues []models.LookupValue
	result := config.DB.Find(&lookupValues)
	return lookupValues, result.Error
}

func CreateLookupValue(lookupValue *models.LookupValue) error {
	result := config.DB.Create(lookupValue)
	return result.Error
}

func GetLookupValueByID(id uint) (*models.LookupValue, error) {
	var lookupValue models.LookupValue
	result := config.DB.First(&lookupValue, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lookupValue, nil
}

func GetLookupValueByKey(key string) (*models.LookupValue, error) {
	var lookupValue models.LookupValue
	result := config.DB.First(&lookupValue, "key = ?", key)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lookupValue, nil
}

func UpdateLookupValue(lookupValue *models.LookupValue) error {
	result := config.DB.Save(lookupValue)
	return result.Error
}

func DeleteLookupValue(id uint) error {
	result := config.DB.Delete(&models.LookupValue{}, id)
	return result.Error
}
