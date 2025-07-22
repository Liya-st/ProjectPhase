package data

import (
	"context"
	"errors"
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{collection: db.Collection("tasks")}
}

func (s *TaskService) CreateTask(task *models.Task) error {

	if task.ID  <= 0{
		return errors.New("invalid task ID")
	}
	var existingTask models.Task
	e := s.collection.FindOne(context.Background(), bson.M{"_id": task.ID}).Decode(&existingTask)
	if e == nil {
		return errors.New("task with that ID already exists")
	}

	if task.Status == "" {
		task.Status = models.Todo
	}
	_, err := s.collection.InsertOne(context.Background(), task)
	return err
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	var task models.Task
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func (s *TaskService) UpdateTask(id int64, task *models.Task) error {
	task.ID = id
	_, err := s.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, task)
	if err != nil {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) DeleteTask(id int64) error {
	result, err := s.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}