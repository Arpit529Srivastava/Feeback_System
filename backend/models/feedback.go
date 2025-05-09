package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Feedback represents a feedback entry
type Feedback struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Message   string             `json:"message" bson:"message"`
	Rating    int                `json:"rating" bson:"rating"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
