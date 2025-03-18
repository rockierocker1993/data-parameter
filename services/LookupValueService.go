package services

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/models"
	"data-parameter/repositories"
	"data-parameter/util"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllLookupValue retrieves all lookup values from the database and returns them as JSON.
//
// @Summary      Get all lookup values
// @Description  This function fetches all lookup values from the database, converts them to DTO format,
//
//	and returns them in the response. If an error occurs during retrieval (except "record not found"),
//	it returns an internal server error.
//
// @Tags         Lookup
// @Produce      json
// @Success      200  {object}  util.Response{data=[]dto.LookupValueDto}  "Successful Response"
// @Failure      500  {object}  util.Response  "Internal Server Error"
func GetAllLookupValue(c *gin.Context) {
	slog.Info("in method GetAllLookupValue")
	lookupValues, err := repositories.GetAllLookupValues()
	var dtoList []dto.LookupValueDto
	if err == nil {
		for _, lookupValue := range lookupValues {
			dtoList = append(dtoList, ToLookupValueDTO(&lookupValue))
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("GetAllLookupValue Failed", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("GetAllLookupValue size %v", string(len(dtoList)))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, dtoList))
}

// CreateLookupValue handles the creation of a new lookup value.
// It validates the request body, checks for duplicate keys, and stores the new record.
//
// @param c *gin.Context - The Gin context containing request and response data.
// @response 400 - If the request body is invalid or the key already exists.
// @response 500 - If an internal server error occurs.
// @response 201 - If the lookup value is successfully created.
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
		baseResponse := util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	}

	//create lookup value
	err := repositories.CreateLookupValue(&lookupValue)
	if err != nil {
		slog.Error("Failed to create lookup value", slog.Any("error", lookupValueDto))
		baseResponse := util.ConstructResponse(c, constant.PRMLV02, constant.Source, nil)
		c.JSON(http.StatusInternalServerError, baseResponse)
	} else {
		slog.Info("Lookup value created successfully")
		lookupValueDto.ID = lookupValue.ID
		baseResponse := util.ConstructResponse(c, constant.Success, constant.General, lookupValueDto)
		c.JSON(http.StatusCreated, baseResponse)
	}
}

// GetLookupValueByID retrieves a lookup value by its ID and returns it as JSON.
//
// @Summary      Get a lookup value by ID
// @Description  This function fetches a lookup value from the database based on the provided ID parameter.
//
//	If the record is not found, it returns a "data not found" response. If another error occurs,
//	it returns an internal server error.
//
// @Tags         Lookup
// @Produce      json
// @Param        id   path      int  true  "Lookup Value ID"
// @Success      200  {object}  util.Response{data=dto.LookupValueDto}  "Successful Response"
// @Failure      400  {object}  util.Response  "Bad Request if id is not found"
// @Failure      500  {object}  util.Response  "Bad Request"
func GetLookupValueByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	lookupValue, err := repositories.GetLookupValueByID(uint(id))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Info("Lookup Value with id %v not found", id)
		baseResponse := util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	} else if err != nil {
		slog.Error("Failed to get lookup value by id", slog.Any("error", err))
		baseResponse := util.ConstructResponse(c, constant.Failed, constant.General, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	}
	slog.Info("GetLookupValueByID successfully")
	baseResponse := util.ConstructResponse(c, constant.Success, constant.General, ToLookupValueDTO(lookupValue))
	c.JSON(http.StatusOK, baseResponse)
}

// GetLookupValueByKey retrieves a lookup value by its Key and returns it as JSON.
//
// @Summary      Get a lookup value by Key
// @Description  This function fetches a lookup value from the database based on the provided Key parameter.
//
//	If the record is not found, it returns a "data not found" response. If another error occurs,
//	it returns an internal server error.
//
// @Tags         Lookup
// @Produce      json
// @Param        key   path      string  true  "Lookup Value Key"
// @Success      200  {object}  util.Response{data=dto.LookupValueDto}  "Successful Response"
// @Failure      400  {object}  util.Response  "Bad Request if the key is not found"
// @Failure      500  {object}  util.Response  "Bad Request"
func GetLookupValueByKey(c *gin.Context) {
	key := c.Param("key")
	lookupValue, err := repositories.GetLookupValueByKey(key)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Info("Lookup Value with key %v not found", key)
		baseResponse := util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	} else if err != nil {
		slog.Error("Failed to get lookup value by key", slog.Any("error", err))
		baseResponse := util.ConstructResponse(c, constant.Failed, constant.General, nil)
		c.JSON(http.StatusInternalServerError, baseResponse)
		return
	}
	slog.Info("GetLookupValueByKey successfully")
	baseResponse := util.ConstructResponse(c, constant.Success, constant.General, ToLookupValueDTO(lookupValue))
	c.JSON(http.StatusOK, baseResponse)
}

func UpdateLookupValue(c *gin.Context) {
	slog.Info("in method UpdateLookupValue")
	id, _ := strconv.Atoi(c.Param("id"))
	var lookupValueDto dto.LookupValueDto
	if err := c.ShouldBindJSON(&lookupValueDto); err != nil {
		baseResponse := util.ConstructResponse(c, constant.Failed, constant.General, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	}

	err := repositories.IsExistLookupValueByID(uint(id))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Info("Lookup Value with id %v not found", id)
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
		return
	} else if err != nil {
		slog.Error("Failed to get lookup value by id %v %v", id, slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	//cek existing key
	err = repositories.IsExistLookupValueByKeyAndIdNot(lookupValueDto.Key, id)
	if err == nil {
		slog.Info("Lookup value with key %v already exist", lookupValueDto.Key)
		baseResponse := util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil)
		c.JSON(http.StatusBadRequest, baseResponse)
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to get lookup value by key %v and id not %v %v", lookupValueDto.Key, id, slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	lookupValueDto.ID = uint(id)
	lookupValue := ToLookupValueModel(&lookupValueDto)
	err = repositories.UpdateLookupValue(&lookupValue)
	if err == nil {
		c.JSON(http.StatusAccepted, util.ConstructResponse(c, constant.Success, constant.General, nil))
		return
	} else {
		slog.Error("UpdateLookupValue Failed", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
}

func DeleteLookupValue(c *gin.Context) {
	slog.Info("in method DeleteLookupValue")
	id, _ := strconv.Atoi(c.Param("id"))
	lookupValue, err := repositories.GetLookupValueByID(uint(id))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Info("DeleteLookupValue Failed, Lookup Value with id %v not found", id)
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
		return
	} else if err != nil {
		slog.Error("DeleteLookupValue Failed, get lookup id", id, slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	lookupValue.BaseModel.DeletedAt = gorm.DeletedAt{Time: time.Now()}
	err = repositories.UpdateLookupValue(lookupValue)
	if err == nil {
		slog.Error("DeleteLookupValue Successfully")
		c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, nil))
		return
	} else {
		slog.Error("DeleteLookupValue Failed, update db %v", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
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
