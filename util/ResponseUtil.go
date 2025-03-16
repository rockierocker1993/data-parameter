package util

import (
	"data-parameter/constant"
	"data-parameter/dto"
	"data-parameter/repositories"
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

	responseMessage, _ := repositories.ResponseMessageFindByCodeAndSorce(code, source)

	return dto.BaseResponse{
		StatusCode: code,
		RequestID:  c.GetHeader(constant.RequestID),
		TitleID:    responseMessage.TitleId,
		TitleEN:    responseMessage.TitleEn,
		DescID:     responseMessage.MessageId,
		DescEN:     responseMessage.MessageEn,
		Source:     source,
		Data:       data,
	}
}
