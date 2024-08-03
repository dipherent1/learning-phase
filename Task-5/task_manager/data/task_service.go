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
	// var task_collection = cliento.Client.Database("taskmgr").Collection("task")
	// t.Count++
	// task.ID = t.Count
	// t.Tasks[task.ID] = task

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
func (t *Taskcollection) GetTaskByID(id int) {
	// task, exists := t.Tasks[id]
	// return task, exists
	return
}

// UpdateTask updates an existing task with new details.
func (t *Taskcollection) UpdateTask(id int, updatedTask models.Task) error {
	// task, exists := t.Tasks[id]
	// if !exists {
	// 	return errors.New("task not found")
	// }

	// if updatedTask.Title != "" {
	// 	task.Title = updatedTask.Title
	// }
	// if updatedTask.Description != "" {
	// 	task.Description = updatedTask.Description
	// }
	// if updatedTask.Priority != "" {
	// 	task.Priority = updatedTask.Priority
	// }
	// if updatedTask.Status != "" {
	// 	task.Status = updatedTask.Status
	// }

	// t.Tasks[id] = task
	return nil
}

// DeleteTask removes a task from the in-memory store.
func (t *Taskcollection) DeleteTask(id int) error {
	// _, exists := t.Tasks[id]
	// if !exists {
	// 	return errors.New("task not found")
	// }
	// delete(t.Tasks, id)
	return nil
}
