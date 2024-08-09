package routers

import (
	"tskmgr/Delivery/controllers"
	infrastructure "tskmgr/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(usercontroller *controllers.Usercontroller, taskcontroller *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", usercontroller.SignupController)
	router.POST("/login", usercontroller.LoginController)

	protected := router.Group("/")
	protected.Use(infrastructure.AuthMiddleware)
	{
		protected.POST("/task", taskcontroller.CreateTask)
		protected.GET("/task/:title", taskcontroller.GetTaskByTitle)
	}

	return router
}
