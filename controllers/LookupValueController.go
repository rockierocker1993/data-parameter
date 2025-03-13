package controllers

import (
	"data-parameter/models"
	"data-parameter/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllLookupValue(c *gin.Context) {
	lookupValues, err := services.GetAllLookupValue()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get lookup value"})
		return
	}
	c.JSON(http.StatusOK, lookupValues)
}

func CreateLookupValue(c *gin.Context) {
	var lookupValue models.LookupValue
	if err := c.ShouldBindJSON(&lookupValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateLookupValue(&lookupValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lookup value"})
		return
	}
	c.JSON(http.StatusCreated, lookupValue)
}

func GetLookupValueByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	lookupValue, err := services.GetLookupValueByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lookup value not found"})
		return
	}
	c.JSON(http.StatusOK, lookupValue)
}

func UpdateLookupValue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	lookupValue, err := services.GetLookupValueByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lookup Value not found"})
		return
	}

	if err := c.ShouldBindJSON(&lookupValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateLookupValue(lookupValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Lookup Value"})
		return
	}
	c.JSON(http.StatusOK, lookupValue)
}

func DeleteLookupValue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteLookupValue(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lookup Value not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Lookup Value deleted"})
}
