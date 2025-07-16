package models

import "time"

type Status string

const (
	StatusNotStarted     Status = "not_started"
	StatusInProgress  Status = "in_progress"
	StatusCompleted   Status = "completed"
)

type Task struct {
	ID      string     `json:"id"`
	Title   string     `json:"title"`
	DueDate *time.Time `json:"due_date,omitempty"`
	Status  *Status    `json:"status,omitempty"`
}
