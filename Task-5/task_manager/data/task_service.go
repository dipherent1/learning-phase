package data

import (
	// "errors"
	// "fmt"
	// "fmt"
	// "tskmgr/config"
	// "tskmgr/config"
	"context"
	"log"
	"tskmgr/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Taskcollection manages tasks in an in-memory store.
type Taskcollection struct {
	Tasks *mongo.Collection // Counter for generating unique task IDs
}

func NewTaskCollection(tasks *mongo.Collection) *Taskcollection {
	return &Taskcollection{Tasks: tasks}
}

func (t *Taskcollection) CreateTask(task models.Task) {
	coll := t.Tasks
	task.Id = primitive.NewObjectID()
	results, err := coll.InsertOne(context.TODO(), task)
	if err != nil {
		panic(err)
	}

	log.Println(results)

}

// ListOfTasks returns a list of all tasks.
func (t *Taskcollection) ListOfTasks() []models.Task {
	coll := t.Tasks
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	// Unpacks the cursor into a slice
	var results []models.Task
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	log.Println(results)
	return results
}

// GetTaskByID returns the task with the specified ID.
func (t *Taskcollection) GetTaskByTitle(title string) (models.Task, error) {
	coll := t.Tasks
	filter := bson.M{"title": title}
	var result models.Task
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	// log.Println(result)
	return result, err
}

// UpdateTask updates an existing task with new details.
func (t *Taskcollection) UpdateTask(title string, updatedTask models.Task) (*mongo.UpdateResult, error) {
	coll := t.Tasks

	// Filter to find the task by title
	filter := bson.M{"title": title}

	// Find the existing task
	var result models.Task
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Update fields only if they are provided in updatedTask
	Title := result.Title
	Description := result.Description
	Priority := result.Priority
	Status := result.Status

	if updatedTask.Title != "" {
		Title = updatedTask.Title
	}

	if updatedTask.Description != "" {
		Description = updatedTask.Description
	}

	if updatedTask.Priority != "" {
		Priority = updatedTask.Priority
	}

	if updatedTask.Status != "" {
		Status = updatedTask.Status
	}

	// Define the update object
	update := bson.M{
		"$set": bson.M{
			"title":       Title,
			"description": Description,
			"priority":    Priority,
			"status":      Status,
		},
	}

	// Perform the update
	updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return updateResult, nil
}


// DeleteTask removes a task from the in-memory store.
func (t *Taskcollection) DeleteTask(title string) error {
	coll := t.Tasks
	filter := bson.M{"title": title}
	_, err := coll.DeleteOne(context.TODO(), filter)

	return err
}
