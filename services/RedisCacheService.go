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
	} else {
		slog.Error("Invalid cache type")
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.Failed, constant.General, nil))
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
