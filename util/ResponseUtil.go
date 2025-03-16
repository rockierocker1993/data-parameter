package util

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/repositories"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse(c *gin.Context, statusCode int, response interface{}) {
	switch res := response.(type) {
	case dto.BaseResponse:
		// Jika langsung berupa struct BaseResponse
		c.JSON(statusCode, res)
	case map[string]interface{}:
		// Jika dalam bentuk map
		c.JSON(statusCode, dto.BaseResponse{
			StatusCode: res["status_code"].(string),
			RequestID:  res["request_id"].(string),
			TitleID:    res["title_id"].(string),
			TitleEN:    res["title_en"].(string),
			DescID:     res["desc_id"].(string),
			DescEN:     res["desc_en"].(string),
			Source:     res["source"].(string),
			Data:       res["data"],
		})
	default:
		// Jika format tidak dikenali, kirim error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid response format"})
	}
}

func ConstructResponse(c *gin.Context, code string, source string, data interface{}) dto.BaseResponse {
	responseMessage, err := repositories.ResponseMessageFindByCodeAndSorce(code, source)
	if err == nil {
		return dto.BaseResponse{
			StatusCode: code,
			RequestID:  c.Value(constant.RequestID).(string),
			TitleID:    responseMessage.TitleId,
			TitleEN:    responseMessage.TitleEn,
			DescID:     responseMessage.MessageId,
			DescEN:     responseMessage.MessageEn,
			Source:     source,
			Data:       data,
		}
	} else {
		slog.Error("Failed to get response message", slog.Any("error", err))
		return dto.BaseResponse{
			StatusCode: code,
			RequestID:  c.Value(constant.RequestID).(string),
			TitleID:    constant.UknownErrorTitleId + fmt.Sprintf(" %v:%v", source, code),
			TitleEN:    constant.UknownErrorTitleId + fmt.Sprintf(" %v:%v", source, code),
			DescID:     constant.UknownErrorDescId,
			DescEN:     constant.UknownErrorDescEn,
			Source:     constant.Source,
			Data:       data,
		}
	}
}
