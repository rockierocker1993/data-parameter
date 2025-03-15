package services

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/models"
	"data-parameter/repositories"
	"data-parameter/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLookupValue() ([]dto.LookupValueDto, error) {
	log.Println("in method GetAllLookupValue")
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
	log.Println("in method CreateLookupValue")

	var lookupValueDto dto.LookupValueDto
	if err := c.ShouldBindJSON(&lookupValueDto); err != nil {
		util.JSONResponse(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("body {}", lookupValueDto)
	lookupValue := ToLookupValueModel(&lookupValueDto)

	//cek existing key
	lookupValueExist, _ := repositories.GetLookupValueByKey(lookupValue.Key)
	if lookupValueExist != nil {
		log.Println("Lookup value with key {} already exist", lookupValue.Key)
		baseResponse := dto.BaseResponse{
			Status:    "failed",
			RequestID: "123",
			TitleID:   "Gagal",
			TitleEN:   "Failed",
			DescID:    "Lookup value dengan key " + lookupValue.Key + " sudah ada",
			DescEN:    "Lookup value with key " + lookupValue.Key + " already exist",
			Source:    constant.Source,
			Data:      nil,
		}
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	}

	//create lookup value
	err := repositories.CreateLookupValue(&lookupValue)
	if err != nil {
		log.Fatal("Failed to create lookup value, error: {}", err)
		baseResponse := dto.BaseResponse{
			Status:    "failed",
			RequestID: "123",
			TitleID:   "Gagal",
			TitleEN:   "Failed",
			DescID:    "Lookup value gagal dibuat",
			DescEN:    "Lookup value failed to create",
			Source:    constant.Source,
			Data:      nil,
		}
		c.JSON(http.StatusInternalServerError, baseResponse)
	} else {
		log.Println("Lookup value created successfully")
		lookupValueDto.ID = lookupValue.ID
		baseResponse := dto.BaseResponse{
			Status:    "success",
			RequestID: "123",
			TitleID:   "Sukses",
			TitleEN:   "Success",
			DescID:    "Lookup value berhasil dibuat",
			DescEN:    "Lookup value created successfully",
			Source:    constant.Source,
			Data:      lookupValueDto,
		}
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
