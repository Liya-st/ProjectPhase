package models

import "time"

type Status string

const (
	Not_started   Status = "Not_started"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
)

type Task struct {
	ID       string    `json:"id" bson:"_id,omitempty"`
	Title    string    `json:"title" bson:"title"`
	DueDate  time.Time `json:"due_date" bson:"due_date"`
	Status   Status    `json:"status" bson:"status"`
}
