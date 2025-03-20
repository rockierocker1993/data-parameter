package services

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/models"
	"data-parameter/repositories"
	"data-parameter/util"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllLookupValue retrieves all lookup values and returns them as JSON.
func GetAllLookupValue(c *gin.Context) {
	slog.Info("in method GetAllLookupValue")
	lookupValues, err := repositories.GetAllLookupValues()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving lookup values", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	dtoList := make([]dto.LookupValueDto, len(lookupValues))
	for i, lookupValue := range lookupValues {
		dtoList[i] = ToLookupValueDTO(&lookupValue)
	}
	slog.Info("Successfully retrieved lookup values", slog.Int("count", len(dtoList)))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, dtoList))
}

// CreateLookupValue handles the creation of a new lookup value.
func CreateLookupValue(c *gin.Context) {
	slog.Info("in method CreateLookupValue")
	var lookupValueDto dto.LookupValueDto

	if valid := util.ValidateRequestSingleField(c, &dto.LookupValueDto{}, constant.LOOKUP_VALUE_PREFIX); !valid {
		slog.Debug("Validation failed")
		return
	}

	slog.Info("Creating new lookup value", slog.String("key", lookupValueDto.Key))

	if exists, _ := repositories.GetLookupValueByKey(lookupValueDto.Key); exists != nil {
		slog.Debug("Lookup value with key already exists", slog.String("key", lookupValueDto.Key))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	lookupValue := ToLookupValueModel(&lookupValueDto)
	if err := repositories.CreateLookupValue(&lookupValue); err != nil {
		slog.Error("Failed to create lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.PRMLV02, constant.Source, nil))
		return
	}

	lookupValueDto.ID = lookupValue.ID
	slog.Info("Lookup value created successfully", slog.Uint64("id", uint64(lookupValue.ID)))
	c.JSON(http.StatusCreated, util.ConstructResponse(c, constant.Success, constant.General, lookupValueDto))
}

// GetLookupValueByID retrieves a lookup value by its ID.
func GetLookupValueByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("in method GetLookupValueByID")
	slog.Info("Retrieving lookup value", slog.Int("id", id))
	lookupValue, err := repositories.GetLookupValueByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("Lookup value not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved lookup value", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToLookupValueDTO(lookupValue)))
}

// GetLookupValueByKey retrieves a lookup value by its key.
func GetLookupValueByKey(c *gin.Context) {
	slog.Info("in method GetLookupValueByKey")
	slog.Info("Retrieving lookup value", slog.Any("key", c.Param("key")))
	lookupValue, err := repositories.GetLookupValueByKey(c.Param("key"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("Lookup value not found", slog.Any("key", c.Param("key")))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved lookup value", slog.Any("key", c.Param("key")))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToLookupValueDTO(lookupValue)))
}

// UpdateLookupValue updates an existing lookup value.
func UpdateLookupValue(c *gin.Context) {
	slog.Info("in method UpdateLookupValue")
	id, _ := strconv.Atoi(c.Param("id"))
	var lookupValueDto dto.LookupValueDto
	if err := c.ShouldBindJSON(&lookupValueDto); err != nil {
		slog.Debug("Invalid request payload", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if valid := util.ValidateRequestSingleField(c, &dto.LookupValueDto{}, constant.LOOKUP_VALUE_PREFIX); !valid {
		slog.Debug("Validation failed")
		return
	}

	_, err := repositories.GetLookupValueByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("Lookup value not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if err := repositories.IsExistLookupValueByKeyAndIdNot(lookupValueDto.Key, id); err == nil {
		slog.Debug("Key already exists", slog.String("key", lookupValueDto.Key))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	lookupValueDto.ID = uint(id)
	lookupValue := ToLookupValueModel(&lookupValueDto)
	if err := repositories.UpdateLookupValue(&lookupValue); err != nil {
		slog.Error("Failed to update lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Lookup value updated successfully", slog.Int("id", id))
	c.JSON(http.StatusAccepted, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// DeleteLookupValue marks a lookup value as deleted.
func DeleteLookupValue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("Deleting lookup value", slog.Int("id", id))
	_, err := repositories.GetLookupValueByID(uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("Lookup value not found", slog.Int("id", id))
			status = http.StatusBadRequest
		}
		slog.Error("Error retrieving lookup value", slog.Any("error", err))
		c.JSON(status, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
		return
	}

	if err := repositories.DeleteLookupValue(uint(id)); err != nil {
		slog.Error("Failed to delete lookup value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Lookup value deleted successfully", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// ToLookupValueDTO converts a LookupValue model to its DTO representation.
func ToLookupValueDTO(lookupValue *models.LookupValue) dto.LookupValueDto {
	return dto.LookupValueDto{
		ID:     lookupValue.ID,
		Key:    lookupValue.Key,
		Value:  lookupValue.Value,
		TextId: lookupValue.TextId,
		TextEn: lookupValue.TextEn,
	}
}

// ToLookupValueModel converts a LookupValueDto to its model representation.
func ToLookupValueModel(lookupValueDto *dto.LookupValueDto) models.LookupValue {
	return models.LookupValue{
		ID:     lookupValueDto.ID,
		Key:    lookupValueDto.Key,
		Value:  lookupValueDto.Value,
		TextId: lookupValueDto.TextId,
		TextEn: lookupValueDto.TextEn,
	}
}
