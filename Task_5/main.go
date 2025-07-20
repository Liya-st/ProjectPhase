package main

import (
	"log"

	"github.com/joho/godotenv"
	"task_manager/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := router.SetupRouter()
	r.Run()
}
