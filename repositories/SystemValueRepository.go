package repositories

import (
	"data-parameter/config"
	"data-parameter/models"
)

func GetAllSystemValues() ([]models.SystemValue, error) {
	var SystemValues []models.SystemValue
	result := config.DB.Find(&SystemValues)
	return SystemValues, result.Error
}

func CreateSystemValue(SystemValue *models.SystemValue) error {
	result := config.DB.Create(SystemValue)
	return result.Error
}

func GetSystemValueByID(id uint) (*models.SystemValue, error) {
	var SystemValue models.SystemValue
	result := config.DB.First(&SystemValue, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SystemValue, nil
}

func GetSystemValueByKey(key string) (*models.SystemValue, error) {
	var SystemValue models.SystemValue
	result := config.DB.First(&SystemValue, "key = ?", key)
	if result.Error != nil {
		return nil, result.Error
	}
	return &SystemValue, nil
}

func IsExistSystemValueByKeyAndIdNot(key string, id int) error {
	var SystemValue models.SystemValue
	result := config.DB.First(&SystemValue, "key = $1 AND ID !=$2", key, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSystemValue(SystemValue *models.SystemValue) error {
	result := config.DB.Save(SystemValue)
	return result.Error
}

func DeleteSystemValue(id uint) error {
	result := config.DB.Delete(&models.SystemValue{}, id)
	return result.Error
}
