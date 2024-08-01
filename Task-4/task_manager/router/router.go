package router

import (
	"tskmgr/controllers"
	"github.com/gin-gonic/gin"
)

// Route initializes the Gin router and defines the API routes.
func Route() *gin.Engine {
	router := gin.Default()

	// Define the API endpoints and their corresponding handlers
	router.GET("/tasks", controllers.ViewTasks)            // Retrieve all tasks
	router.POST("/tasks", controllers.CreateTask)          // Create a new task
	router.GET("/tasks/:id", controllers.GetTaskByID)      // Retrieve a task by ID
	router.PUT("/tasks/:id", controllers.UpdateTask)       // Update a task by ID
	router.DELETE("/tasks/:id", controllers.DeleteTask)    // Delete a task by ID

	return router
}
