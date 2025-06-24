package main

import (
	"log"
	"net/http"
	"time"

	"testovoe/api"
	"testovoe/internal/service"
	"testovoe/internal/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	taskStorage := storage.NewMemoryTaskStorage()

	taskService := service.NewTaskService(taskStorage)

	router := gin.Default()

	api.SetupTaskRoutes(router, taskService)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
