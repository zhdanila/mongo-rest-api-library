package main

import (
	"context"
	"fmt"
	"log"
	"mongo"
	"mongo/internal/handler"
	"mongo/internal/repository"
	"mongo/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	client, err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	repo := repository.NewRepository(client)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(mongo.Server)

	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			log.Fatalf("error with running server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	fmt.Println("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error with shutting down server: %s", err.Error())
	}
}
