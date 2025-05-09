package routes

import (
	"feedback-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupFeedbackRoutes(router *gin.Engine, feedbackController *controllers.FeedbackController) {
	feedbackRoutes := router.Group("/api/feedback")
	{
		feedbackRoutes.POST("/", feedbackController.CreateFeedback)
		feedbackRoutes.GET("/", feedbackController.GetAllFeedback)
	}
}
