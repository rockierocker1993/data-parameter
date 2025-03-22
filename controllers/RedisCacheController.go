package controllers

import (
	"data-parameter/services"

	"github.com/gin-gonic/gin"
)

func ReloadCache(c *gin.Context) {
	services.ReloadCache(c)
}
