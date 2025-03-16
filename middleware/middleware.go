package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware menambahkan Request ID ke setiap request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID") // Cek apakah client mengirim Request ID

		if requestID == "" {
			requestID = uuid.New().String() // Generate UUID jika tidak ada
		}

		// Set Request ID ke context Gin
		c.Set("RequestID", requestID)

		// Tambahkan Request ID ke log
		log.Printf("[Request ID: %s] %s %s", requestID, c.Request.Method, c.Request.URL.Path)

		// Tambahkan header ke response
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()
	}
}
