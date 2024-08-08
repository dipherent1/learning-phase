package main

import (
	"context"
	"log"
	"tskmgr/Delivery/config"
	"tskmgr/Delivery/routers"
)

func main() {
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
