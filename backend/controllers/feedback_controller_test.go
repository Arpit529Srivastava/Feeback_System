package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"feedback-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection is a mock implementation of mongo.Collection
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

// MockCursor is a mock implementation of mongo.Cursor
type MockCursor struct {
	mock.Mock
}

func (m *MockCursor) All(ctx context.Context, result interface{}) error {
	args := m.Called(ctx, result)
	return args.Error(0)
}

func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestCreateFeedback(t *testing.T) {
	// Setup
	router := setupTestRouter()
	mockCollection := new(MockCollection)
	controller := &FeedbackController{collection: mockCollection}

	// Test data
	feedback := models.Feedback{
		Name:    "Test User",
		Email:   "test@example.com",
		Message: "Test feedback message",
		Rating:  5,
	}

	// Mock InsertOne
	insertedID := primitive.NewObjectID()
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(
		&mongo.InsertOneResult{InsertedID: insertedID},
		nil,
	)

	// Setup route
	router.POST("/feedback", controller.CreateFeedback)

	// Test cases
	t.Run("successful feedback creation", func(t *testing.T) {
		// Create request
		jsonData, _ := json.Marshal(feedback)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusCreated, w.Code)

		var response models.Feedback
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, feedback.Name, response.Name)
		assert.Equal(t, feedback.Email, response.Email)
		assert.Equal(t, feedback.Message, response.Message)
		assert.Equal(t, feedback.Rating, response.Rating)
		assert.NotEmpty(t, response.ID)
		assert.NotEmpty(t, response.CreatedAt)
	})

	t.Run("missing required fields", func(t *testing.T) {
		testCases := []struct {
			name     string
			feedback models.Feedback
			status   int
		}{
			{
				name: "missing name",
				feedback: models.Feedback{
					Email:   "test@example.com",
					Message: "Test message",
					Rating:  5,
				},
				status: http.StatusBadRequest,
			},
			{
				name: "missing email",
				feedback: models.Feedback{
					Name:    "Test User",
					Message: "Test message",
					Rating:  5,
				},
				status: http.StatusBadRequest,
			},
			{
				name: "missing message",
				feedback: models.Feedback{
					Name:   "Test User",
					Email:  "test@example.com",
					Rating: 5,
				},
				status: http.StatusBadRequest,
			},
			{
				name: "missing rating",
				feedback: models.Feedback{
					Name:    "Test User",
					Email:   "test@example.com",
					Message: "Test message",
				},
				status: http.StatusBadRequest,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				jsonData, _ := json.Marshal(tc.feedback)
				req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()

				router.ServeHTTP(w, req)
				assert.Equal(t, tc.status, w.Code)
			})
		}
	})

	t.Run("invalid feedback data", func(t *testing.T) {
		// Create request with invalid data
		invalidData := []byte(`{"invalid": "data"`)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(invalidData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("database error", func(t *testing.T) {
		// Reset mock expectations
		mockCollection.ExpectedCalls = nil

		// Mock database error
		mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(
			nil,
			mongo.ErrNoDocuments,
		).Once()

		jsonData, _ := json.Marshal(feedback)
		req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("rating validation", func(t *testing.T) {
		testCases := []struct {
			name          string
			rating        int
			expectedCode  int
			shouldPass    bool
			errorContains string
		}{
			{
				name:         "valid rating 1",
				rating:       1,
				expectedCode: http.StatusCreated,
				shouldPass:   true,
			},
			{
				name:         "valid rating 3",
				rating:       3,
				expectedCode: http.StatusCreated,
				shouldPass:   true,
			},
			{
				name:         "valid rating 5",
				rating:       5,
				expectedCode: http.StatusCreated,
				shouldPass:   true,
			},
			{
				name:          "invalid rating 0",
				rating:        0,
				expectedCode:  http.StatusOK,
				shouldPass:    false,
				errorContains: "rating must be between 1 and 5",
			},
			{
				name:          "invalid rating 6",
				rating:        6,
				expectedCode:  http.StatusBadRequest,
				shouldPass:    false,
				errorContains: "rating must be between 1 and 5",
			},
			{
				name:          "invalid rating -1",
				rating:        -1,
				expectedCode:  http.StatusBadRequest,
				shouldPass:    false,
				errorContains: "rating must be between 1 and 5",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Reset mock expectations
				mockCollection.ExpectedCalls = nil

				// Create feedback with test rating
				testFeedback := models.Feedback{
					Name:    "Test User",
					Email:   "test@example.com",
					Message: "Test message",
					Rating:  tc.rating,
				}

				// Set up mock for valid ratings
				if tc.shouldPass {
					insertedID := primitive.NewObjectID()
					mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(
						&mongo.InsertOneResult{InsertedID: insertedID},
						nil,
					).Once()
				}

				jsonData, _ := json.Marshal(testFeedback)
				req, _ := http.NewRequest("POST", "/feedback", bytes.NewBuffer(jsonData))
				req.Header.Set("Content-Type", "application/json")
				w := httptest.NewRecorder()

				router.ServeHTTP(w, req)

				// Check status code
				assert.Equal(t, tc.expectedCode, w.Code)

				if tc.shouldPass {
					// For valid ratings, verify the response
					var response models.Feedback
					err := json.Unmarshal(w.Body.Bytes(), &response)
					assert.NoError(t, err)
					assert.Equal(t, tc.rating, response.Rating)
					assert.NotEmpty(t, response.ID)
					assert.NotEmpty(t, response.CreatedAt)
				} else {
					// For invalid ratings, verify the error message
					var response map[string]interface{}
					err := json.Unmarshal(w.Body.Bytes(), &response)
					assert.NoError(t, err)
					assert.Contains(t, response["error"], tc.errorContains)
				}
			})
		}
	})
}

func TestGetAllFeedback(t *testing.T) {
	// Setup
	router := setupTestRouter()
	mockCollection := new(MockCollection)
	controller := &FeedbackController{collection: mockCollection}

	// Test data
	now := time.Now()
	feedbacks := []models.Feedback{
		{
			ID:        primitive.NewObjectID(),
			Name:      "User 1",
			Email:     "user1@example.com",
			Message:   "Feedback 1",
			Rating:    4,
			CreatedAt: now,
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "User 2",
			Email:     "user2@example.com",
			Message:   "Feedback 2",
			Rating:    5,
			CreatedAt: now,
		},
	}

	// Setup route
	router.GET("/feedback", controller.GetAllFeedback)

	// Test cases
	t.Run("successful get all feedback", func(t *testing.T) {
		// Mock Find
		mockCollection.On("Find", mock.Anything, bson.M{}).Return(&mongo.Cursor{}, nil).Once()

		// Create request
		req, _ := http.NewRequest("GET", "/feedback", nil)
		w := httptest.NewRecorder()

		// Perform request
		controller.decodeCursor = func(cursor *mongo.Cursor, out interface{}) error {
			feedbackPtr := out.(*[]models.Feedback)
			*feedbackPtr = feedbacks
			return nil
		}
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Feedback
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)

		// Validate all fields for first feedback
		assert.Equal(t, feedbacks[0].ID, response[0].ID)
		assert.Equal(t, feedbacks[0].Name, response[0].Name)
		assert.Equal(t, feedbacks[0].Email, response[0].Email)
		assert.Equal(t, feedbacks[0].Message, response[0].Message)
		assert.Equal(t, feedbacks[0].Rating, response[0].Rating)
		assert.GreaterOrEqual(t, response[0].Rating, 1)
		assert.LessOrEqual(t, response[0].Rating, 5)
		assert.Equal(t, feedbacks[0].CreatedAt.Unix(), response[0].CreatedAt.Unix())

		// Validate all fields for second feedback
		assert.Equal(t, feedbacks[1].ID, response[1].ID)
		assert.Equal(t, feedbacks[1].Name, response[1].Name)
		assert.Equal(t, feedbacks[1].Email, response[1].Email)
		assert.Equal(t, feedbacks[1].Message, response[1].Message)
		assert.Equal(t, feedbacks[1].Rating, response[1].Rating)
		assert.GreaterOrEqual(t, response[1].Rating, 1)
		assert.LessOrEqual(t, response[1].Rating, 5)
		assert.Equal(t, feedbacks[1].CreatedAt.Unix(), response[1].CreatedAt.Unix())
	})

	t.Run("empty feedback list", func(t *testing.T) {
		// Mock Find
		mockCollection.On("Find", mock.Anything, bson.M{}).Return(&mongo.Cursor{}, nil).Once()

		// Create request
		req, _ := http.NewRequest("GET", "/feedback", nil)
		w := httptest.NewRecorder()

		// Perform request with empty feedback list
		controller.decodeCursor = func(cursor *mongo.Cursor, out interface{}) error {
			feedbackPtr := out.(*[]models.Feedback)
			*feedbackPtr = []models.Feedback{}
			return nil
		}
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		var response []models.Feedback
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})

	t.Run("database error", func(t *testing.T) {
		// Mock Find with error
		mockCollection.On("Find", mock.Anything, bson.M{}).Return(nil, mongo.ErrNoDocuments).Once()

		// Create request
		req, _ := http.NewRequest("GET", "/feedback", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
