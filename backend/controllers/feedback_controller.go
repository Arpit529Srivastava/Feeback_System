package controllers

import (
	"context"
	"net/http"
	"time"

	"feedback-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeedbackController struct {
	collection *mongo.Collection
}

func NewFeedbackController(db *mongo.Database) *FeedbackController {
	return &FeedbackController{
		collection: db.Collection("feedbacks"),
	}
}

func (fc *FeedbackController) CreateFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback.CreatedAt = time.Now()
	result, err := fc.collection.InsertOne(context.Background(), feedback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feedback"})
		return
	}

	feedback.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, feedback)
}

func (fc *FeedbackController) GetAllFeedback(c *gin.Context) {
	var feedbacks []models.Feedback
	cursor, err := fc.collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feedbacks"})
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &feedbacks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode feedbacks"})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}
