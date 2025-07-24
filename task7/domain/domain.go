package domain

import (
	"context"
	"errors"
	"time"
	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type TaskStatus string

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

var (
	ErrInvalidTaskID     = errors.New("invalid task ID")
	ErrInvalidTaskTitle  = errors.New("task title cannot be empty")
	ErrInvalidTaskStatus = errors.New("invalid task status")
	ErrInvalidDueDate    = errors.New("due date cannot be in the past")
	ErrInvalidUsername   = errors.New("username cannot be empty")
	ErrInvalidRole       = errors.New("invalid user role")
	ErrUsernameExists    = errors.New("username already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrAdminRequired     = errors.New("admin access required")
)

type Task struct {
	ID      string     `json:"id"`
	Title   string     `json:"title"`
	Status  TaskStatus `json:"status"`
	DueDate time.Time  `json:"due_date"`
}

func NewTask(title string, status TaskStatus, dueDate time.Time) (*Task, error) {
	if title == "" {
		return nil, ErrInvalidTaskTitle
	}
	if status != Todo && status != InProgress && status != Done {
		return nil, ErrInvalidTaskStatus
	}

	if status == "" {
		status = Todo
	}
	return &Task{
		ID:      uuid.New().String(),
		Title:   title,
		Status:  status,
		DueDate: dueDate,
	}, nil
}

func (t *Task) UpdateStatus(newStatus TaskStatus) error {
	if newStatus != Todo && newStatus != InProgress && newStatus != Done {
		return ErrInvalidTaskStatus
	}
	t.Status = newStatus
	return nil
}

type User struct {
	ID           string 
	Username     string 
	PasswordHash string  
	Role         Role   
}

func NewUser(username, passwordHash string, role Role) (*User, error) {
	if username == "" {
		return nil, ErrInvalidUsername
	}
	if passwordHash == "" {
		return nil, ErrInvalidPassword
	}
	if role != RoleAdmin && role != RoleUser {
		return nil, ErrInvalidRole
	}
	return &User{
		ID:           uuid.New().String(),
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
	}, nil
}

type TaskRepository interface {
	Create(c context.Context, task *Task) (*Task, error)
	GetByID(c context.Context, id string) (*Task, error)
	Update(c context.Context, task *Task) error
	Delete(c context.Context, id string) error
	GetAll(c context.Context) ([]*Task, error)
}

type UserRepository interface {
	Create(c context.Context, user *User) (*User, error)
	GetByID(c context.Context, id string) (*User, error)
	GetByUsername(c context.Context, username string) (*User, error)
	Update(c context.Context, user *User) error
	CountUsers(c context.Context) (int64, error)
	}