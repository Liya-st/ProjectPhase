package router

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService(db)
	userService := data.NewUserService(db)
	controller := controllers.NewController(taskService, userService)

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(userService))
	{
		protected.GET("/tasks", controller.GetAllTasks)
		protected.GET("/tasks/:id", controller.GetTaskByID)

		admin := protected.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", controller.CreateTask)
			admin.PUT("/tasks/:id", controller.UpdateTask)
			admin.DELETE("/tasks/:id", controller.DeleteTask)
			admin.PUT("/promote/:username", controller.PromoteUser)
		}
	}

	return r
}