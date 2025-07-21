package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	StatusNotStarted Status = "not_started"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `json:"title" bson:"title" validate:"required"`
	DueDate   *time.Time         `json:"due_date,omitempty" bson:"due_date,omitempty"`
	Status    Status             `json:"status" bson:"status" validate:"oneof=not_started in_progress completed"`
	// UserID    primitive.ObjectID `json:"-" bson:"user_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

