package main

import (
	"log"
	"os"
	"subs-manager-api/internal/handlers"
	"subs-manager-api/internal/server"
	"subs-manager-api/internal/services/subscribe"
	"subs-manager-api/internal/storage/postgresql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	initEnv()

	// TODO initialize logger

	storage := postgresql.NewPostgres(os.Getenv("DATABASE_URL"))

	subService := subscribe.NewSubscribeService(storage)

	handlers := handlers.NewHandler(subService)

	r := gin.Default()

	srv := server.NewServer(handlers, r)

	srv.Start(os.Getenv("PORT"))
}

func initEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("No .env file found or failed to load: %v", err)
	}
}
