package services

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/models"
	"data-parameter/repositories"
	"data-parameter/util"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLookupValue() ([]dto.LookupValueDto, error) {
	slog.Info("in method GetAllLookupValue")
	lookupValues, err := repositories.GetAllLookupValues()
	var dtoList []dto.LookupValueDto
	if err == nil {
		for _, lookupValue := range lookupValues {
			dtoList = append(dtoList, ToLookupValueDTO(&lookupValue))
		}
	}
	return dtoList, err
}

func CreateLookupValue(c *gin.Context) {
	slog.Info("in method CreateLookupValue")

	var lookupValueDto dto.LookupValueDto
	if err := c.ShouldBindJSON(&lookupValueDto); err != nil {
		util.JSONResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Info("Request Body", slog.Any("body", lookupValueDto))
	lookupValue := ToLookupValueModel(&lookupValueDto)

	//cek existing key
	lookupValueExist, _ := repositories.GetLookupValueByKey(lookupValue.Key)
	if lookupValueExist != nil {
		log.Printf("Lookup value with key %v already exist", lookupValue.Key)
		baseResponse := util.ConstructResponse(c, "PRMLV01", constant.Source, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	}

	//create lookup value
	err := repositories.CreateLookupValue(&lookupValue)
	if err != nil {
		slog.Error("Failed to create lookup value", slog.Any("error", lookupValueDto))
		baseResponse := util.ConstructResponse(c, "PRMLV02", constant.Source, nil)
		c.JSON(http.StatusInternalServerError, baseResponse)
	} else {
		slog.Info("Lookup value created successfully")
		lookupValueDto.ID = lookupValue.ID
		baseResponse := util.ConstructResponse(c, constant.Success, constant.General, lookupValueDto)
		c.JSON(http.StatusBadRequest, baseResponse)
	}
}

func GetLookupValueByID(id uint) (*models.LookupValue, error) {
	return repositories.GetLookupValueByID(id)
}

func GetLookupValueByKey(key string) (*models.LookupValue, error) {
	return repositories.GetLookupValueByKey(key)
}

func UpdateLookupValue(lookupValue *models.LookupValue) error {
	return repositories.UpdateLookupValue(lookupValue)
}

func DeleteLookupValue(id uint) error {
	return repositories.DeleteLookupValue(id)
}

func ToLookupValueDTO(lookupValue *models.LookupValue) dto.LookupValueDto {
	return dto.LookupValueDto{
		ID:     lookupValue.ID,
		Key:    lookupValue.Key,
		Value:  lookupValue.Value,
		TextId: lookupValue.TextId,
		TextEn: lookupValue.TextEn,
	}
}

func ToLookupValueModel(lookupValueDto *dto.LookupValueDto) models.LookupValue {
	return models.LookupValue{
		ID:     lookupValueDto.ID,
		Key:    lookupValueDto.Key,
		Value:  lookupValueDto.Value,
		TextId: lookupValueDto.TextId,
		TextEn: lookupValueDto.TextEn,
	}
}
