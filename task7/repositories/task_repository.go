package repositories

import (
	"context"
	"task_manager/domain"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo")

type TaskRepositoryImpl struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) domain.TaskRepository {
	return &TaskRepositoryImpl{collection: db.Collection("tasks")}
}

func (r *TaskRepositoryImpl) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
	_, err := r.collection.InsertOne(c, bson.M{
		"_id":      task.ID,
		"title":    task.Title,
		"status":   task.Status,
		"due_date": task.DueDate,
	})
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepositoryImpl) GetByID(c context.Context, id string) (*domain.Task, error) {
	var result struct {
		ID      string           `bson:"_id" json:"id"`
		Title   string           `bson:"title" json:"title"`
		Status  domain.TaskStatus `bson:"status" json:"status"`
		DueDate time.Time        `bson:"due_date" json:"due_date"`
	}
	err := r.collection.FindOne(c, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrInvalidTaskID
		}
		return nil, err
	}
	return &domain.Task{
		ID:      result.ID,
		Title:   result.Title,
		Status:  result.Status,
		DueDate: result.DueDate,
	}, nil
}

func (r *TaskRepositoryImpl) Update(c context.Context, task *domain.Task) error {
	result, err := r.collection.UpdateOne(c, bson.M{"_id": task.ID}, bson.M{
		"$set": bson.M{
			"title":    task.Title,
			"status":   task.Status,
			"due_date": task.DueDate,
		},
	})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return domain.ErrInvalidTaskID
	}
	return nil
}

func (r *TaskRepositoryImpl) Delete(c context.Context, id string) error {
	result, err := r.collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return domain.ErrInvalidTaskID
	}
	return nil
}

func (r *TaskRepositoryImpl) GetAll(c context.Context) ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor, err := r.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var result struct {
			ID      string           `bson:"_id" json:"id"`
			Title   string           `bson:"title" json:"title"`
			Status  domain.TaskStatus `bson:"status" json:"status"`
			DueDate time.Time        `bson:"due_date" json:"due_date"`
	}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		tasks = append(tasks, &domain.Task{
			ID:      result.ID,
			Title:   result.Title,
			Status:  result.Status,
			DueDate: result.DueDate,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}