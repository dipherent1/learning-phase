package routers

import (
	"tskmgr/Delivery/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine{
	router := gin.Default()
	
	usercoll:= client.Database("TaskManager").Collection("Users")
	
	mycontroller := controllers.NewController(usercoll)

	router.POST("/signup", mycontroller.SignupController)

	return router
}
