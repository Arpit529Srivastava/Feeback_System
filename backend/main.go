package main

import (
	"feedback-backend/config"
	"feedback-backend/controllers"
	"feedback-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	feedbackController := controllers.NewFeedbackController(config.DB.Collection("feedbacks"))
	routes.SetupFeedbackRoutes(r, feedbackController)
	routes.SetupHealthRoutes(r)
	r.Run(":8080")
	// removed the cors policy because it was not working..done
}
