package main

import (
	"context"
	"log"
	"tskmgr/Delivery/config"
	"tskmgr/Delivery/routers"

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

	router := routers.SetupRouter(client)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}

}
