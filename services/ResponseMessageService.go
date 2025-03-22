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

// GetAllResponseMessage retrieves all response messages and returns them as JSON.
func GetAllResponseMessage(c *gin.Context) {
	slog.Info("in method GetAllResponseMessage")
	ResponseMessages, err := repositories.GetAllResponseMessages()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Error retrieving response messages", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	dtoList := make([]dto.ResponseMessageDto, len(ResponseMessages))
	for i, ResponseMessage := range ResponseMessages {
		dtoList[i] = ToResponseMessageDTO(&ResponseMessage)
	}
	slog.Info("Successfully retrieved response messages", slog.Int("count", len(dtoList)))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, dtoList))
}

// CreateResponseMessage handles the creation of a new response message.
func CreateResponseMessage(c *gin.Context) {
	slog.Info("in method CreateResponseMessage")
	var ResponseMessageDto dto.ResponseMessageDto

	if valid := util.ValidateRequestSingleField(c, &dto.ResponseMessageDto{}, constant.LOOKUP_VALUE); !valid {
		slog.Debug("Validation failed")
		return
	}

	slog.Info("Creating new response message", slog.String("code", ResponseMessageDto.Code))

	if exists, _ := repositories.GetResponseMessageByCode(ResponseMessageDto.Code); exists != nil {
		slog.Debug("response message with code already exists", slog.String("code", ResponseMessageDto.Code))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	ResponseMessage := ToResponseMessageModel(&ResponseMessageDto)
	if err := repositories.CreateResponseMessage(&ResponseMessage); err != nil {
		slog.Error("Failed to create response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.PRMLV02, constant.Source, nil))
		return
	}

	ResponseMessageDto.ID = ResponseMessage.ID
	slog.Info("response message created successfully", slog.Uint64("id", uint64(ResponseMessage.ID)))
	c.JSON(http.StatusCreated, util.ConstructResponse(c, constant.Success, constant.General, ResponseMessageDto))
}

// GetResponseMessageByID retrieves a response message by its ID.
func GetResponseMessageByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("in method GetResponseMessageByID")
	slog.Info("Retrieving response message", slog.Int("id", id))
	ResponseMessage, err := repositories.GetResponseMessageByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("response message not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved response message", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToResponseMessageDTO(ResponseMessage)))
}

// GetResponseMessageByCode retrieves a response message by its code.
func GetResponseMessageByCode(c *gin.Context) {
	slog.Info("in method GetResponseMessageByCode")
	slog.Info("Retrieving response message", slog.Any("code", c.Param("code")))
	ResponseMessage, err := repositories.GetResponseMessageByCode(c.Param("code"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Warn("response message not found", slog.Any("code", c.Param("code")))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("Successfully retrieved response message", slog.Any("code", c.Param("code")))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, ToResponseMessageDTO(ResponseMessage)))
}

// UpdateResponseMessage updates an existing response message.
func UpdateResponseMessage(c *gin.Context) {
	slog.Info("in method UpdateResponseMessage")
	id, _ := strconv.Atoi(c.Param("id"))
	var ResponseMessageDto dto.ResponseMessageDto
	if err := c.ShouldBindJSON(&ResponseMessageDto); err != nil {
		slog.Debug("Invalid request payload", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if valid := util.ValidateRequestSingleField(c, &dto.ResponseMessageDto{}, constant.LOOKUP_VALUE); !valid {
		slog.Debug("Validation failed")
		return
	}

	_, err := repositories.GetResponseMessageByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("response message not found", slog.Int("id", id))
			c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
			return
		}
		slog.Error("Error retrieving response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}

	if err := repositories.IsExistResponseMessageByCodeAndIdNot(ResponseMessageDto.Code, id); err == nil {
		slog.Debug("Code already exists", slog.String("code", ResponseMessageDto.Code))
		c.JSON(http.StatusBadRequest, util.ConstructResponse(c, constant.PRMLV01, constant.Source, nil))
		return
	}

	ResponseMessageDto.ID = uint(id)
	ResponseMessage := ToResponseMessageModel(&ResponseMessageDto)
	if err := repositories.UpdateResponseMessage(&ResponseMessage); err != nil {
		slog.Error("Failed to update response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("response message updated successfully", slog.Int("id", id))
	c.JSON(http.StatusAccepted, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// DeleteResponseMessage marks a response message as deleted.
func DeleteResponseMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	slog.Info("Deleting response message", slog.Int("id", id))
	_, err := repositories.GetResponseMessageByID(uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Debug("response message not found", slog.Int("id", id))
			status = http.StatusBadRequest
		}
		slog.Error("Error retrieving response message", slog.Any("error", err))
		c.JSON(status, util.ConstructResponse(c, constant.DATA_NOT_FOUND, constant.General, nil))
		return
	}

	if err := repositories.DeleteResponseMessage(uint(id)); err != nil {
		slog.Error("Failed to delete response message", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, util.ConstructResponse(c, constant.Failed, constant.General, nil))
		return
	}
	slog.Info("response message deleted successfully", slog.Int("id", id))
	c.JSON(http.StatusOK, util.ConstructResponse(c, constant.Success, constant.General, nil))
}

// ToResponseMessageDTO converts a ResponseMessage model to its DTO representation.
func ToResponseMessageDTO(ResponseMessage *models.ResponseMessage) dto.ResponseMessageDto {
	return dto.ResponseMessageDto{
		ID:        ResponseMessage.ID,
		Code:      ResponseMessage.Code,
		TitleId:   ResponseMessage.TitleId,
		TitleEn:   ResponseMessage.TitleEn,
		MessageId: ResponseMessage.MessageId,
		MessageEn: ResponseMessage.MessageEn,
		Source:    ResponseMessage.Source,
	}
}

// ToResponseMessageModel converts a ResponseMessageDto to its model representation.
func ToResponseMessageModel(ResponseMessageDto *dto.ResponseMessageDto) models.ResponseMessage {
	return models.ResponseMessage{
		ID:        ResponseMessageDto.ID,
		Code:      ResponseMessageDto.Code,
		TitleId:   ResponseMessageDto.TitleId,
		TitleEn:   ResponseMessageDto.TitleEn,
		MessageId: ResponseMessageDto.MessageId,
		MessageEn: ResponseMessageDto.MessageEn,
		Source:    ResponseMessageDto.Source,
	}
}
