package controllers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"feedback-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FeedbackCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

type FeedbackController struct {
	collection   FeedbackCollection
	decodeCursor func(cursor *mongo.Cursor, out interface{}) error
}

func NewFeedbackController(collection FeedbackCollection) *FeedbackController {
	return &FeedbackController{
		collection: collection,
		decodeCursor: func(cursor *mongo.Cursor, out interface{}) error {
			return cursor.All(context.Background(), out)
		},
	}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

// validateFeedback checks if all required fields are present and valid
func validateFeedback(feedback models.Feedback) error {
	if feedback.Name == "" {
		return ValidationError{Field: "name", Message: "name is required"}
	}
	if feedback.Email == "" {
		return ValidationError{Field: "email", Message: "email is required"}
	}
	if feedback.Message == "" {
		return ValidationError{Field: "message", Message: "message is required"}
	}
	if feedback.Rating < 1 || feedback.Rating > 5 {
		return ValidationError{Field: "rating", Message: "rating must be between 1 and 5"}
	}
	return nil
}

func (fc *FeedbackController) CreateFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate feedback
	if err := validateFeedback(feedback); err != nil {
		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationErr.Message,
				"field": validationErr.Field,
			})
			return
		}
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
	c.Header("Content-Type", "application/json")
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

	if err = fc.decodeCursor(cursor, &feedbacks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode feedbacks"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, feedbacks)
}
