package controllers

import (
	"data-parameter/services"

	"github.com/gin-gonic/gin"
)

func GetAllSystemValue(c *gin.Context) {
	services.GetAllSystemValue(c)
}

func CreateSystemValue(c *gin.Context) {
	services.CreateSystemValue(c)
}

func GetSystemValueByID(c *gin.Context) {
	services.GetSystemValueByID(c)
}

func GetSystemValueByKey(c *gin.Context) {
	services.GetSystemValueByKey(c)
}

func UpdateSystemValue(c *gin.Context) {
	services.UpdateSystemValue(c)
}

func DeleteSystemValue(c *gin.Context) {
	services.DeleteSystemValue(c)
}
