package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/core/services"
	"github.com/samuelorlato/task-manager-api/internal/handlers"
)

func main() {
	engine := gin.Default()

	taskService := services.NewTaskService()

	HTTPHandler := handlers.NewHTTPHandler(engine, taskService)
	HTTPHandler.SetRoutes()

	engine.Run()
}
