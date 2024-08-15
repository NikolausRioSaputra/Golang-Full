package routes

import (
	"consume-api/internal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the Gin router with all the necessary routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/rest", handler.ConsumeRest)
	r.POST("/soap", handler.ConsumeSOAP)

	return r
}
