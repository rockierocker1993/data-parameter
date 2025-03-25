package services

import (
	"context"
	"data-parameter/config"
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/models"
	"data-parameter/repositories"
	"data-parameter/util"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReloadCache(c *gin.Context) {
	cacheType := c.Query("cacheType")
	module := c.Query("module")
	code := c.Query("code")
	key := c.Query("key")

	slog.Info("in method ReloadCache table %v cacheModule %v code %v key %v", cacheType, module, code, key)
	if cacheType == constant.LOOKUP_VALUE {
		reloadCacheLookupValue(key)

	} else if cacheType == constant.ALL {
		reloadCacheLookupValue("")
		reloadSystemValues()

	} else if cacheType == constant.SYSTEM_VALUE {
		if module != "" && key != "" {
			reloadSystemValuesByModuleAndKey(module, key)
		} else if module != "" {
			reloadSystemValuesByModule(module)
		} else {
			reloadSystemValues()
		}

	} else {
		slog.Error("Invalid cache type")
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.INVALID_CACHE_TYPE, constant.General, nil))
		return

	}

	slog.Info("Successfully reloaded cache")
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

func reloadCacheLookupValue(lookupKey string) error {
	slog.Info("in method reloadCacheLookupValue")
	ctx := context.Background()
	lookupValues := make([]models.LookupValue, 0)
	err := error(nil)
	if lookupKey != "" {
		lookupValues, err = repositories.GetAllLookupValuesByKey(lookupKey)
	} else {
		lookupValues, err = repositories.GetAllLookupValues()
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving lookup values", slog.Any("error", err))
		return err
	}
	dtoList := make([]dto.LookupValueDto, len(lookupValues))
	for i, lookupValue := range lookupValues {
		dtoList[i] = ToLookupValueDTO(&lookupValue)
	}

	groupedLookupValues := groupLookupValueByKey(dtoList)
	for key, value := range groupedLookupValues {
		objString, _ := util.ObjectToString(value)
		key := constant.LOOKUP_VALUE + "_" + key
		config.RDB.Set(ctx, key, objString, 0)
	}
	return nil
}

func groupLookupValueByKey(lookupValues []dto.LookupValueDto) map[string]dto.LookupValueDto {
	slog.Info("in method groupLookupValueByKey")
	groupedLookupValues := make(map[string]dto.LookupValueDto)
	for _, lookupValue := range lookupValues {
		groupedLookupValues[lookupValue.Key] = lookupValue
	}
	return groupedLookupValues
}

func reloadSystemValuesByModuleAndKey(module string, key string) error {
	slog.Info("in method reloadSystemValuesByModuleAndKey")
	ctx := context.Background()
	systemValue, err := repositories.GetSystemValueByModuleAndKey(module, key)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving system value", slog.Any("error", err))
		return err
	}
	systemValueDto := ToSystemValueDTO(systemValue)
	objString, _ := util.ObjectToString(systemValueDto)
	key = constant.SYSTEM_VALUE + "_" + module + "_" + key
	config.RDB.Set(ctx, key, objString, 0)
	return nil
}

func reloadSystemValuesByModule(module string) error {
	slog.Info("in method reloadSystemValuesByModule")
	ctx := context.Background()
	systemValues, err := repositories.GetSystemValueByModule(module)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving system values", slog.Any("error", err))
		return err
	}
	for i := 0; i < len(systemValues); i++ {
		systemValueDto := ToSystemValueDTO(&systemValues[i])
		key := constant.SYSTEM_VALUE + "_" + systemValueDto.Module + "_" + systemValueDto.Key
		objString, _ := util.ObjectToString(systemValueDto)
		config.RDB.Set(ctx, key, objString, 0)
	}
	return nil
}

func reloadSystemValues() error {
	slog.Info("in method reloadSystemValues")
	ctx := context.Background()
	systemValues, err := repositories.GetAllSystemValues()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving system values", slog.Any("error", err))
		return err
	}
	for i := 0; i < len(systemValues); i++ {
		systemValueDto := ToSystemValueDTO(&systemValues[i])
		key := constant.SYSTEM_VALUE + "_" + systemValueDto.Module + "_" + systemValueDto.Key
		objString, _ := util.ObjectToString(systemValueDto)
		config.RDB.Set(ctx, key, objString, 0)
	}
	return nil
}

func reloadCacheResponseMessage(code string) error {
	slog.Info("in method reloadCacheLookupValue")
	ctx := context.Background()
	responseMessages := make([]models.ResponseMessage, 0)
	err := error(nil)
	if code != "" {
		var responseMessage *models.ResponseMessage
		responseMessage, err = repositories.GetResponseMessageByCode(code)
		responseMessages[0] = models.ResponseMessage{
			ID:        responseMessage.ID,
			Code:      responseMessage.Code,
			MessageId: responseMessage.MessageId,
			MessageEn: responseMessage.MessageEn,
			Source:    responseMessage.Source,
		}
	} else {
		responseMessages, err = repositories.GetAllResponseMessages()
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving response messages", slog.Any("error", err))
		return err
	}
	for i := 0; i < len(responseMessages); i++ {
		responseMessageDto := ToResponseMessageDTO(&responseMessages[i])
		objString, _ := util.ObjectToString(responseMessageDto)
		key := constant.RESPONSE_MESSAGE + "_" + responseMessageDto.Code
		config.RDB.Set(ctx, key, objString, 0)
	}
	return nil
}
