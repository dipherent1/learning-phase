package router

import (
	"tskmgr/controllers"

	"github.com/gin-gonic/gin"
)

// Route initializes the Gin router and defines the API routes.
func Route(tc *controllers.Taskcontroller) *gin.Engine {

	router := gin.Default()

	// Define the API endpoints and their corresponding handlers
	router.GET("/tasks", tc.ViewTasks)         // Retrieve all tasks
	router.POST("/tasks", tc.CreateTask)       // Create a new task
	router.GET("/tasks/:title", tc.GetTaskByID)   // Retrieve a task by ID
	router.PUT("/tasks/:title", tc.UpdateTask)    // Update a task by ID
	router.DELETE("/tasks/:title", tc.DeleteTask) // Delete a task by ID

	return router
}
