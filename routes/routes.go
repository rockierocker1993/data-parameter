package routes

import (
	"data-parameter/config"
	"data-parameter/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Gunakan middleware Request ID
	r.Use(config.RequestIDMiddleware())
	contextPath := os.Getenv("CONTEXT_PATH")
	api := r.Group(contextPath)
	{
		//lookup value api
		api.GET("/lookup-value", controllers.GetAllLookupValue)
		api.POST("/lookup-value", controllers.CreateLookupValue)
		api.GET("/lookup-value/:id", controllers.GetLookupValueByID)
		api.GET("/lookup-value/key/:key", controllers.GetLookupValueByKey)
		api.PUT("/lookup-value/:id", controllers.UpdateLookupValue)
		api.DELETE("/lookup-value/:id", controllers.DeleteLookupValue)

		//system value api
		api.GET("/system-value", controllers.GetAllSystemValue)
		api.POST("/system-value", controllers.CreateSystemValue)
		api.GET("/system-value/:id", controllers.GetSystemValueByID)
		api.GET("/system-value/module/key/:module/:key", controllers.GetSystemValueByKey)
		api.PUT("/system-value/:id", controllers.UpdateSystemValue)
		api.DELETE("/system-value/:id", controllers.DeleteSystemValue)

		//response message api
		api.GET("/response-message", controllers.GetAllResponseMessage)
		api.POST("/response-message", controllers.CreateResponseMessage)
		api.GET("/response-message/:id", controllers.GetResponseMessageByID)
		api.GET("/response-message/code/:code", controllers.GetResponseMessageByCode)
		api.PUT("/response-message/:id", controllers.UpdateResponseMessage)
		api.DELETE("/response-message/:id", controllers.DeleteResponseMessage)
	}
	return r
}
