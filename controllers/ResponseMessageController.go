package controllers

import (
	"data-parameter/services"

	"github.com/gin-gonic/gin"
)

func GetAllResponseMessage(c *gin.Context) {
	services.GetAllResponseMessage(c)
}

func CreateResponseMessage(c *gin.Context) {
	services.CreateResponseMessage(c)
}

func GetResponseMessageByID(c *gin.Context) {
	services.GetResponseMessageByID(c)
}

func GetResponseMessageByCode(c *gin.Context) {
	services.GetResponseMessageByCode(c)
}

func UpdateResponseMessage(c *gin.Context) {
	services.UpdateResponseMessage(c)
}

func DeleteResponseMessage(c *gin.Context) {
	services.DeleteResponseMessage(c)
}
