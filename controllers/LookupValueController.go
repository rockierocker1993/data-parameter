package controllers

import (
	"data-parameter/services"

	"github.com/gin-gonic/gin"
)

func GetAllLookupValue(c *gin.Context) {
	services.GetAllLookupValue(c)
}

func CreateLookupValue(c *gin.Context) {
	services.CreateLookupValue(c)
}

func GetLookupValueByID(c *gin.Context) {
	services.GetLookupValueByID(c)
}

func GetLookupValueByKey(c *gin.Context) {
	services.GetLookupValueByKey(c)
}

func UpdateLookupValue(c *gin.Context) {
	services.UpdateLookupValue(c)
}

func DeleteLookupValue(c *gin.Context) {
	services.DeleteLookupValue(c)
}
