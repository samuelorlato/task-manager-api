package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuelorlato/task-manager-api/internal/configs"
	"github.com/samuelorlato/task-manager-api/internal/core/services"
	"github.com/samuelorlato/task-manager-api/internal/handlers"
	"github.com/samuelorlato/task-manager-api/internal/repositories"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	credentialsJSONString := os.Getenv("FIREBASE_CREDENTIALS")
	app, err := configs.InitFirebaseApp(ctx, credentialsJSONString)
	if err != nil {
		panic(err)
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	firestoreRepository := repositories.NewFirestoreRepository(firestoreClient)

	taskService := services.NewTaskService(firestoreRepository)

	bcryptService := services.NewBcryptService()
	userService := services.NewUserService(firestoreRepository, bcryptService)

	jwtService := services.NewJWTService()

	errorHandler := handlers.NewErrorHandler()

	HTTPHandler := handlers.NewHTTPHandler(engine, taskService, userService, bcryptService, jwtService, errorHandler)
	HTTPHandler.SetRoutes()

	engine.Run()
}
