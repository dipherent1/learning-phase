package main

import (
	"context"
	"log"
	"tskmgr/Delivery/config"
	"tskmgr/Delivery/routers"
	repositories "tskmgr/Repositories"
	usecases "tskmgr/Usecases"
	controllers "tskmgr/Delivery/controllers"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	client := config.ConnectDB()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	
	usercoll := client.Database("TaskManager").Collection("Users")
	taskcoll := client.Database("TaskManager").Collection("Tasks")

	// generate user repository
	userrepo := repositories.NewUserDataManipulator(usercoll)

	// generate task repository
	taskrepo := repositories.NewTaskDataManipulator(taskcoll)

	// generate user usecase
	userusecase := usecases.NewUserUsecase(userrepo)

	//gerate task usecase
	taskusecase := usecases.NewTaskUsecase(taskrepo)

	usercontroller := controllers.NewUsercontroller(userusecase)
	taskcontroller := controllers.NewTaskController(taskusecase)

	router := routers.SetupRouter(usercontroller, taskcontroller)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}

}
