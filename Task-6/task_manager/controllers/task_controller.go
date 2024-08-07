package controllers

import (
	"net/http"
	"tskmgr/data"
	"tskmgr/models"

	"github.com/gin-gonic/gin"
)

// TaskController handles HTTP requests for tasks
type TaskController struct {
	taskService *data.TaskService
}

// NewTaskController initializes a new TaskController
func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{taskService: service}
}

// ViewTasks handles GET /tasks to retrieve all tasks
func (tc *TaskController) ViewTasks(c *gin.Context) {

	tasks, err := tc.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// CreateTask handles POST /tasks to create a new task
func (tc *TaskController) CreateTask(c *gin.Context) {
	claim, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claim not set"})
		return
	}

	userClaims, ok := claim.(models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		return
	}

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	userid := userClaims.UserId
	newTask, err := tc.taskService.CreateTask(newTask, userid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// GetTaskByID handles GET /tasks/:id to retrieve a task by ID
func (tc *TaskController) GetTaskByTitle(c *gin.Context) {
	title := c.Param("title")

	task, err := tc.taskService.GetTaskByTitle(title)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask handles PUT /tasks/:id to update a task by ID
func (tc *TaskController) UpdateTask(c *gin.Context) {
	claim, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claim not set"})
		return
	}

	// Type assertion to your custom Claims type
	userClaims, ok := claim.(models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		return
	}

	if userClaims.UserRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only admins can update tasks"})
		return
	}

	title := c.Param("title")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	updateResult, err := tc.taskService.UpdateTask(title, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResult)
}

// DeleteTask handles DELETE /tasks/:id to delete a task by ID
func (tc *TaskController) DeleteTask(c *gin.Context) {
	claim, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claim not set"})
		return
	}

	// Type assertion to your custom Claims type
	userClaims, ok := claim.(models.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		return
	}

	if userClaims.UserRole != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "only admins can delete tasks"})
		return
	}

	title := c.Param("title")

	err := tc.taskService.DeleteTask(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
