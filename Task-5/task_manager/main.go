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
	// Initialize MongoDB connection
	client := config.ConnectDB()
	if client == nil {
		log.Fatal("Failed to initialize MongoDB client")
	}

	// Create a new task collection service
	taskService := data.NewTaskService(client.Database("TaskManager").Collection("Tasks"))

	// Initialize task controller with the task service
	taskController := controllers.NewTaskController(taskService)

	// Ensure the MongoDB client disconnects when the application exits
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Set up router and start the server
	r := router.SetupRouter(taskController)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
