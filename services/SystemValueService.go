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

// GetAllSystemValue retrieves all system values and returns them as JSON.
func GetAllSystemValue(c *gin.Context) {
	slog.Info("in method GetAllSystemValue")
	SystemValues, err := repositories.GetAllSystemValues()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving system values", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	dtoList := make([]dto.SystemValueDto, len(SystemValues))
	for i, SystemValue := range SystemValues {
		dtoList[i] = ToSystemValueDTO(&SystemValue)
	}
	slog.Info("Successfully retrieved system values", slog.Int("count", len(dtoList)))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, dtoList))
}

// CreateSystemValue handles the creation of a new System value.
func CreateSystemValue(c *gin.Context) {
	slog.Info("in method CreateSystemValue")
	var SystemValueDto dto.SystemValueDto

	if valid := util.ValidateRequestSingleField(c, &dto.SystemValueDto{}, constant.SYSTEM_VALUE_PREFIX); !valid {
		slog.Debug("Validation failed")
		return
	}

	slog.Info("Creating new System value", slog.String("key", SystemValueDto.Key))

	if exists, _ := repositories.GetSystemValueByModuleAndKey(SystemValueDto.Module, SystemValueDto.Key); exists != nil {
		slog.Debug("System value with module %v and key %v already exists", slog.String("key", SystemValueDto.Key), slog.String("key", SystemValueDto.Key))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	SystemValue := ToSystemValueModel(&SystemValueDto)
	if err := repositories.CreateSystemValue(&SystemValue); err != nil {
		slog.Error("Failed to create System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.PRMLV02, constant.Source, nil))
		return
	}

	SystemValueDto.ID = SystemValue.ID
	slog.Info("System value created successfully", slog.Uint64("id", uint64(SystemValue.ID)))
	c.JSON(http.StatusCreated, util.ConstructResponse(c, constant.Success, constant.General, SystemValueDto))
}

// GetSystemValueByID retrieves a System value by its ID.
func GetSystemValueByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("in method GetSystemValueByID")
	slog.Info("Retrieving System value", slog.Int("id", id))
	SystemValue, err := repositories.GetSystemValueByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("System value not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved System value", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToSystemValueDTO(SystemValue)))
}

// GetSystemValueByKey retrieves a System value by its key.
func GetSystemValueByKey(c *gin.Context) {
	module := c.Param("module")
	key := c.Param("key")
	slog.Info("in method GetSystemValueByKey")
	slog.Info("Retrieving System value module %v key %v", module, key)
	SystemValue, err := repositories.GetSystemValueByModuleAndKey(module, key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("System value not found module %v key %v", module, key)
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved System value module %v key %v", module, key)
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToSystemValueDTO(SystemValue)))
}

// UpdateSystemValue updates an existing System value.
func UpdateSystemValue(c *gin.Context) {
	slog.Info("in method UpdateSystemValue")
	id, _ := strconv.Atoi(c.Param("id"))
	var SystemValueDto dto.SystemValueDto
	if err := c.ShouldBindJSON(&SystemValueDto); err != nil {
		slog.Debug("Invalid request payload", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if valid := util.ValidateRequestSingleField(c, &dto.SystemValueDto{}, constant.SYSTEM_VALUE_PREFIX); !valid {
		slog.Debug("Validation failed")
		return
	}

	_, err := repositories.GetSystemValueByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("System value not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if err := repositories.IsExistSystemValueByModuleAndKeyAndIdNot(SystemValueDto.Module, SystemValueDto.Key, id); err == nil {
		slog.Debug("module %v key %v already exists", SystemValueDto.Module, SystemValueDto.Key)
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	SystemValueDto.ID = uint(id)
	SystemValue := ToSystemValueModel(&SystemValueDto)
	if err := repositories.UpdateSystemValue(&SystemValue); err != nil {
		slog.Error("Failed to update System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("System value updated successfully", slog.Int("id", id))
	c.JSON(http.StatusAccepted, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// DeleteSystemValue marks a System value as deleted.
func DeleteSystemValue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("Deleting System value", slog.Int("id", id))
	_, err := repositories.GetSystemValueByID(uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("System value not found", slog.Int("id", id))
			status = http.StatusBadRequest
		}
		slog.Error("Error retrieving System value", slog.Any("error", err))
		c.JSON(status, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
		return
	}

	if err := repositories.DeleteSystemValue(uint(id)); err != nil {
		slog.Error("Failed to delete System value", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("System value deleted successfully", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// ToSystemValueDTO converts a SystemValue model to its DTO representation.
func ToSystemValueDTO(SystemValue *models.SystemValue) dto.SystemValueDto {
	return dto.SystemValueDto{
		ID:        SystemValue.ID,
		Module:    SystemValue.Module,
		Key:       SystemValue.Key,
		Value:     SystemValue.Value,
		IsEncrypt: SystemValue.IsEncrypt,
	}
}

// ToSystemValueModel converts a SystemValueDto to its model representation.
func ToSystemValueModel(SystemValueDto *dto.SystemValueDto) models.SystemValue {
	return models.SystemValue{
		ID:        SystemValueDto.ID,
		Module:    SystemValueDto.Module,
		Key:       SystemValueDto.Key,
		Value:     SystemValueDto.Value,
		IsEncrypt: SystemValueDto.IsEncrypt,
	}
}
