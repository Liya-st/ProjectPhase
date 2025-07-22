package models

type TaskStatus string

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

type Task struct {
	ID          int64     `json:"id" bson:"_id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Status      TaskStatus `json:"status" bson:"status"`
}