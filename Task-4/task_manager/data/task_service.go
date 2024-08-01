package data

import (
	"errors"
	"tskmgr/models"
)

// TaskManager manages tasks in an in-memory store.
type TaskManager struct {
	Tasks map[int]models.Task // In-memory task storage
	Count int                 // Counter for generating unique task IDs
}

// NewTaskManager creates a new TaskManager instance.
func NewTaskManager() *TaskManager {
	return &TaskManager{
		Tasks: make(map[int]models.Task),
		Count: 0,
	}
}

// CreateTask adds a new task to the in-memory store.
func (t *TaskManager) CreateTask(task models.Task) {
	t.Count++
	task.ID = t.Count
	t.Tasks[task.ID] = task
}

// ListOfTasks returns a list of all tasks.
func (t *TaskManager) ListOfTasks() []models.Task {
	var tasks []models.Task
	for _, task := range t.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskByID returns the task with the specified ID.
func (t *TaskManager) GetTaskByID(id int) (models.Task, bool) {
	task, exists := t.Tasks[id]
	return task, exists
}

// UpdateTask updates an existing task with new details.
func (t *TaskManager) UpdateTask(id int, updatedTask models.Task) error {
	task, exists := t.Tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}
	if updatedTask.Priority != "" {
		task.Priority = updatedTask.Priority
	}
	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}

	t.Tasks[id] = task
	return nil
}

// DeleteTask removes a task from the in-memory store.
func (t *TaskManager) DeleteTask(id int) error {
	_, exists := t.Tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(t.Tasks, id)
	return nil
}
