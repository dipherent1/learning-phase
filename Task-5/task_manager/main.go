package main

import (
	"context"

	"log"
	"tskmgr/config"
	"tskmgr/controllers"
	"tskmgr/data"

	"tskmgr/router"
)

func main() {

	client := config.ConnectDB()

	if client == nil {
		log.Fatal("client is not initialized")
	}

	tasks := data.NewTaskCollection(client.Database("Taskmgr").Collection("Tasks"))

	taskcontroller := controllers.NewTaskController(*tasks)

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	r := router.Route(taskcontroller)
	r.Run(":8080")
	log.Println("I am here end")
}
