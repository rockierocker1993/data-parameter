package services

import (
	"data-parameter/models"
	"data-parameter/repositories"
)

func GetAllLookupValue() ([]models.LookupValue, error) {
	return repositories.GetAllLookupValues()
}

func CreateLookupValue(lookupValue *models.LookupValue) error {
	return repositories.CreateLookupValue(lookupValue)
}

func GetLookupValueByID(id uint) (*models.LookupValue, error) {
	return repositories.GetLookupValueByID(id)
}

func UpdateLookupValue(lookupValue *models.LookupValue) error {
	return repositories.UpdateLookupValue(lookupValue)
}

func DeleteLookupValue(id uint) error {
	return repositories.DeleteLookupValue(id)
}
