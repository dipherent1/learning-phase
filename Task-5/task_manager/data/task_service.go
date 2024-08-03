package data

import (
	"context"
	"log"
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
func (ts *TaskService) CreateTask(task models.Task) {
	task.Id = primitive.NewObjectID()
	_, err := ts.taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatalf("Failed to insert task: %v", err)
	}
}

// GetAllTasks retrieves all tasks from the MongoDB collection
func (ts *TaskService) GetAllTasks() []models.Task {
	cursor, err := ts.taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("Failed to find tasks: %v", err)
	}
	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		log.Fatalf("Failed to decode tasks: %v", err)
	}
	return tasks
}

// GetTaskByTitle retrieves a task by its title
func (ts *TaskService) GetTaskByTitle(title string) (models.Task, error) {
	var task models.Task
	err := ts.taskCollection.FindOne(context.TODO(), bson.M{"title": title}).Decode(&task)
	return task, err
}

// UpdateTask updates a task in the MongoDB collection
func (ts *TaskService) UpdateTask(title string, updatedTask models.Task) (*mongo.UpdateResult, error) {
	filter := bson.M{"title": title}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"priority":    updatedTask.Priority,
			"status":      updatedTask.Status,
		},
	}
	return ts.taskCollection.UpdateOne(context.TODO(), filter, update)
}

// DeleteTask removes a task from the MongoDB collection
func (ts *TaskService) DeleteTask(title string) error {
	_, err := ts.taskCollection.DeleteOne(context.TODO(), bson.M{"title": title})
	return err
}
