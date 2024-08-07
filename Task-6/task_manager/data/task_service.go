package data

import (
	"context"
	// "log"
	"tskmgr/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskService provides task management services
type TaskService struct {
	taskCollection *mongo.Collection
}

// NewTaskService initializes a new TaskService
func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{taskCollection: collection}
}

// CreateTask inserts a new task into the MongoDB collection
func (ts *TaskService) CreateTask(task models.Task) error {
	task.Id = primitive.NewObjectID()
	_, err := ts.taskCollection.InsertOne(context.TODO(), task)

	if err != nil {
		return err
	}
	return nil
}

// GetAllTasks retrieves all tasks from the MongoDB collection
func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	cursor, err := ts.taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetTaskByTitle retrieves a task by its title
func (ts *TaskService) GetTaskByTitle(title string) (models.Task, error) {
	var task models.Task
	filter := bson.M{"title": title}
	err := ts.taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	return task, err
}

// UpdateTask updates a task in the MongoDB collection
func (ts *TaskService) UpdateTask(title string, updatedTask models.Task) (models.Task, error) {
	task, err := ts.GetTaskByTitle(title)
	if err != nil {
		return models.Task{}, err
	}

	var Title, Description, Priority, Status string
	
	if updatedTask.Title != "" {
		Title = updatedTask.Title
	} else {
		Title = task.Title

	}

	if updatedTask.Description != "" {
		Description = updatedTask.Description
	} else {
		Description = task.Description
	}

	if updatedTask.Priority != "" {
		Priority = updatedTask.Priority

	} else {
		Priority = task.Priority
	}

	if updatedTask.Status != "" {
		Status = updatedTask.Status

	} else {
		Status = task.Status
	}

	// Define the update object
	updatedData := bson.M{
		"$set": bson.M{
			"title":       Title,
			"description": Description,
			"priority":    Priority,
			"status":      Status,
		},
	}

	filter := bson.M{"title": title}

	_, err = ts.taskCollection.UpdateOne(context.TODO(), filter, updatedData)

	if err != nil {
		return models.Task{}, err
	}
	
	updatedTaskData, err := ts.GetTaskByTitle(Title)
	if err != nil {
		return models.Task{}, err
	}

	return updatedTaskData, nil
}

// DeleteTask removes a task from the MongoDB collection
func (ts *TaskService) DeleteTask(title string) error {
	_, err := ts.taskCollection.DeleteOne(context.TODO(), bson.M{"title": title})
	return err
}
