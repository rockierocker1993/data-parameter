package routes

import (
	"data-parameter/controllers"
	"data-parameter/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Gunakan middleware Request ID
	r.Use(middleware.RequestIDMiddleware())
	constextPath := os.Getenv("CONTEXT_PATH")
	api := r.Group(constextPath)
	{
		//lookup value api
		api.GET("/lookup-value", controllers.GetAllLookupValue)
		api.POST("/lookup-value", controllers.CreateLookupValue)
		api.GET("/lookup-value/:id", controllers.GetLookupValueByID)
		api.PUT("/lookup-value/:id", controllers.UpdateLookupValue)
		api.DELETE("/lookup-value/:id", controllers.DeleteLookupValue)
	}
	return r
}
