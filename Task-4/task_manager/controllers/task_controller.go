package controllers

import (
	"net/http"
	"strconv"
	"tskmgr/data"
	"tskmgr/models"

	"github.com/gin-gonic/gin"
)

// Task manager instance to handle business logic and data manipulation
var taskManager = data.NewTaskManager()

// ViewTasks handles GET /tasks, returns a list of all tasks.
func ViewTasks(c *gin.Context) {
	tasks := taskManager.ListOfTasks()
	c.JSON(http.StatusOK, tasks)
}

// CreateTask handles POST /tasks, creates a new task.
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	taskManager.CreateTask(newTask)
	c.JSON(http.StatusCreated, newTask)
}

// GetTaskByID handles GET /tasks/:id, returns the task with the specified ID.
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, exists := taskManager.GetTaskByID(taskID)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask handles PUT /tasks/:id, updates the task with the specified ID.
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	err = taskManager.UpdateTask(taskID, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles DELETE /tasks/:id, deletes the task with the specified ID.
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = taskManager.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
