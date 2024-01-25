package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	handler "testTask/pkg/handlers"
	"testTask/pkg/service"
	"testTask/repository"
	"testTask/server"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	log.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("Error occured on db connection close: %s", err.Error())
	}
}
