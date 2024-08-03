package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Task represents a task with an ID, title, description, priority, and status.
type Task struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"` 
	Title       string `json:"title"`       // Title of the task
	Description string `json:"description"` // Description of the task
	Priority    string `json:"priority"`    // Priority level of the task
	Status      string `json:"status"`      // Current status of the task (e.g., "pending", "completed")
}
