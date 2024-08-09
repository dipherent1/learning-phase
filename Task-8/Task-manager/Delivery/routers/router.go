package routers

import (
	"tskmgr/Delivery/controllers"
	infrastructure "tskmgr/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine{
	router := gin.Default()
	
	usercoll:= client.Database("TaskManager").Collection("Users")
	taskcoll:= client.Database("TaskManager").Collection("Tasks")
	
	usercontroller := controllers.NewUserController(usercoll)
	taskcontroller := controllers.NewTaskController(taskcoll)

	router.POST("/signup", usercontroller.SignupController)
	router.POST("/login", usercontroller.LoginController)

	protected:=router.Group("/")
	protected.Use(infrastructure.AuthMiddleware)
	{
		protected.POST("/task",taskcontroller.CreateTask)
		protected.GET("/task/:title",taskcontroller.GetTaskByTitle)
	}
	
	return router
}
