package main

import (
	"log"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task_manager/data"
	"task_manager/router"
)

func main() {
	// MongoDB connection
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	if err = client.Ping(nil, nil); err != nil {
		log.Fatal("MongoDB not responding:", err)
	}

	// Init collections
	data.InitUserCollection(client)
	data.InitTaskCollection(client)
	log.Println("✅ Connected to MongoDB")

	

	// Load all routes from router file
	r := router.SetupRouter()

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
