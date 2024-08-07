package router

import (
	"tskmgr/controllers"
	"tskmgr/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and defines the API routes
func SetupRouter(tc *controllers.TaskController, uc *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", uc.SignupController)
	router.POST("/login", uc.LoginController)

	// Define the protected routes
	router.GET("/tasks", tc.ViewTasks)             // Retrieve all tasks
	router.GET("/tasks/:title", tc.GetTaskByTitle) // Retrieve a task by title

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware)
	{
		// Use protected instead of router to ensure middleware is applied
		protected.POST("/tasks", tc.CreateTask)          // Create a new task
		protected.PUT("/tasks/:title", tc.UpdateTask)    // Update a task by title
		protected.DELETE("/tasks/:title", tc.DeleteTask) // Delete a task by title

		protected.GET("/users", middleware.IsSuperUser, uc.GetUsers)
		protected.PUT("/users/:username", middleware.IsSuperUser, uc.ChangeRole)

	}

	return router
}
