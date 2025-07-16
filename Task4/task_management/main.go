package main

import (
	"task_management/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
