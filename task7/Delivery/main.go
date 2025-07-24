package main

import (
	"context"
	"log"
	"os"
	"task_manager/Delivery/controllers"
	"task_manager/infrastructure"
	"task_manager/repositories"
	router "task_manager/Delivery/routers"
	"task_manager/usecases"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	db := client.Database(os.Getenv("MONGO_DB"))

	taskRepo := repositories.NewTaskRepository(db)
	userRepo := repositories.NewUserRepository(db)

	passwordService := infrastructure.NewPasswordService()
	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"))

	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, passwordService, jwtService)

	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	r := router.SetupRouter(taskController, userController, jwtService)
	r.Run(":8080")
}
