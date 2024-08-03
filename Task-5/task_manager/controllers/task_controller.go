package controllers

import (
	// "log"
	"net/http"
	"tskmgr/data"
	"tskmgr/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Taskcontroller struct {
	coll data.Taskcollection
}

func NewTaskController(cur_col data.Taskcollection) *Taskcontroller {
	return &Taskcontroller{
		coll: cur_col,
	}
}

func (tc *Taskcontroller) ViewTasks(c *gin.Context) {
	tasks := tc.coll.ListOfTasks()
	c.JSON(http.StatusOK, tasks)
}

// CreateTask handles POST /tasks, creates a new task.
func (tc *Taskcontroller) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}
	
	newTask.Id = primitive.NewObjectID()

	tc.coll.CreateTask(newTask)
	c.JSON(http.StatusCreated, newTask)
}

// GetTaskByID handles GET /tasks/:id, returns the task with the specified ID.
func (tc *Taskcontroller) GetTaskByID(c *gin.Context) {
	title := c.Param("title")

	task, err := tc.coll.GetTaskByTitle(title)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask handles PUT /tasks/:id, updates the task with the specified ID.
func (tc *Taskcontroller) UpdateTask(c *gin.Context) {
	title := c.Param("title")
	

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided"})
		return
	}

	updateResult, err := tc.coll.UpdateTask(title, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResult)
}

// DeleteTask handles DELETE /tasks/:id, deletes the task with the specified ID.
func (tc *Taskcontroller) DeleteTask(c *gin.Context) {
	title := c.Param("title")

	err := tc.coll.DeleteTask(title)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
