package data

import (
	"errors"
	"task_management/models"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Learn Go", Status: strPtr(models.StatusNotStarted), DueDate: timePtr(time.Now().AddDate(0, 0, 7))},
	{ID: "2", Title: "Build a web service", Status: strPtr(models.StatusNotStarted), DueDate: timePtr(time.Now().AddDate(0, 0, 14))},
}

func strPtr(s models.Status) *models.Status {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task models.Task) {
	tasks = append(tasks, task)
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func UpdateTask(id string, updated models.Task) error {
	for i, t := range tasks {
		if t.ID == id {
			if updated.Title != "" {
				tasks[i].Title = updated.Title
			}
			if updated.DueDate != nil {
				tasks[i].DueDate = updated.DueDate
			}
			if updated.Status != nil {
				tasks[i].Status = updated.Status
			}
			return nil
		}
	}
	return errors.New("task not found")
}
