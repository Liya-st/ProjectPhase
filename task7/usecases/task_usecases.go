package usecases

import (
	"context"
	"task_manager/domain"
	"time"
)

type TaskInput struct {
	Title   string
	Status  domain.TaskStatus
	DueDate time.Time
}

type TaskUsecase struct {
	repo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) CreateTask(c context.Context, input TaskInput) (*domain.Task, error) {
	task, err := domain.NewTask(input.Title, input.Status, input.DueDate)
	if err != nil {
		return nil, err
	}
	task, err = u.repo.Create(c, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *TaskUsecase) GetTask(c context.Context, id string) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrInvalidTaskID
	}
	task, err := u.repo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *TaskUsecase) UpdateTask(c context.Context, id string, input TaskInput) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrInvalidTaskID
	}
	task, err := u.repo.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	updatedTask, err := domain.NewTask(input.Title, input.Status, input.DueDate)
	if err != nil {
		return nil, err
	}
	updatedTask.ID = task.ID
	err = u.repo.Update(c, updatedTask)
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

func (u *TaskUsecase) DeleteTask(c context.Context, id string) error {
	if id == "" {
		return domain.ErrInvalidTaskID
	}
	return u.repo.Delete(c, id)
}

func (u *TaskUsecase) ListTasks(c context.Context) ([]*domain.Task, error) {
	tasks, err := u.repo.GetAll(c)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
