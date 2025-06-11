package routes

import (
	"feedback-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)
}
