package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/core/services"
	"github.com/samuelorlato/task-manager-api/internal/handlers"
	"github.com/samuelorlato/task-manager-api/internal/repositories"
)

func main() {
	engine := gin.Default()

	taskRepository := repositories.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository)

	HTTPHandler := handlers.NewHTTPHandler(engine, taskService)
	HTTPHandler.SetRoutes()

	engine.Run()
}
