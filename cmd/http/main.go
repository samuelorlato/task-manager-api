package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/samuelorlato/task-manager-api/internal/configs"
	"github.com/samuelorlato/task-manager-api/internal/core/services"
	"github.com/samuelorlato/task-manager-api/internal/handlers"
	"github.com/samuelorlato/task-manager-api/internal/repositories"
)

func main() {
	ctx := context.Background()

	app, err := configs.InitFirebaseApp(ctx)
	if err != nil {
		panic(err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	taskRepository := repositories.NewFirestoreTaskRepository(firestoreClient)
	taskService := services.NewTaskService(taskRepository)

	errorHandler := handlers.NewErrorHandler()

	HTTPHandler := handlers.NewHTTPHandler(engine, taskService, errorHandler)
	HTTPHandler.SetRoutes()

	engine.Run()
}
