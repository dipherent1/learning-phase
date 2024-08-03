package controllers

import (
	// "log"
	"net/http"
	"strconv"
	"tskmgr/data"
	"tskmgr/models"

	"github.com/gin-gonic/gin"
)

type Taskcontroller struct {
	service data.Taskcollection // Counter for generating unique task IDs
}

func NewTaskController(serv data.Taskcollection) *Taskcontroller {
	return &Taskcontroller{
		service: serv,
	}
}

func (tc *Taskcontroller) ViewTasks(c *gin.Context) {
	tasks := tc.service.ListOfTasks()
	c.JSON(http.StatusOK, tasks)
}

// CreateTask handles POST /tasks, creates a new task.
func (tc *Taskcontroller) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	tc.service.CreateTask(newTask)
	c.JSON(http.StatusCreated, newTask)
}

// GetTaskByID handles GET /tasks/:id, returns the task with the specified ID.
func (tc *Taskcontroller) GetTaskByID(c *gin.Context) {
	// id := c.Param("id")
	// taskID, err := strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
	// 	return
	// }

	// // task, exists := tc.service.GetTaskByID(taskID)
	// if !exists {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	// 	return
	// }

	// c.JSON(http.StatusOK, task)
}

// UpdateTask handles PUT /tasks/:id, updates the task with the specified ID.
func (tc *Taskcontroller) UpdateTask(c *gin.Context) {
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

	err = tc.service.UpdateTask(taskID, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles DELETE /tasks/:id, deletes the task with the specified ID.
func (tc *Taskcontroller) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = tc.service.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
