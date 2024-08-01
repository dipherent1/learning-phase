package models

// Task represents a task with an ID, title, description, priority, and status.
type Task struct {
	ID          int    `json:"id"`          // Unique identifier for the task
	Title       string `json:"title"`       // Title of the task
	Description string `json:"description"` // Description of the task
	Priority    string `json:"priority"`    // Priority level of the task
	Status      string `json:"status"`      // Current status of the task (e.g., "pending", "completed")
}
