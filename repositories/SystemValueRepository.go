package repositories

import (
	"data-parameter/config"
	"data-parameter/models"
)

func GetAllSystemValues() ([]models.SystemValue, error) {
	var systemValues []models.SystemValue
	result := config.DB.Find(&systemValues)
	return systemValues, result.Error
}

func CreateSystemValue(systemValue *models.SystemValue) error {
	result := config.DB.Create(systemValue)
	return result.Error
}

func GetSystemValueByID(id uint) (*models.SystemValue, error) {
	var systemValue models.SystemValue
	result := config.DB.First(&systemValue, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &systemValue, nil
}

func GetSystemValueByModuleAndKey(module string, key string) (*models.SystemValue, error) {
	var systemValue models.SystemValue
	result := config.DB.First(&systemValue, "module = $1 AND key = $2", module, key)
	if result.Error != nil {
		return nil, result.Error
	}
	return &systemValue, nil
}

func IsExistSystemValueByModuleAndKeyAndIdNot(module string, key string, id int) error {
	var systemValue models.SystemValue
	result := config.DB.First(&systemValue, "module = $1 AND key = $2 AND ID !=$3", key, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSystemValue(systemValue *models.SystemValue) error {
	result := config.DB.Save(systemValue)
	return result.Error
}

func DeleteSystemValue(id uint) error {
	result := config.DB.Delete(&models.SystemValue{}, id)
	return result.Error
}
