package router

import (
	"task_manager/Delivery/controllers"
	"task_manager/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController, jwtService infrastructure.JWTService) *gin.Engine {
	r := gin.Default()
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	protected := r.Group("/api")
	protected.Use(infrastructure.AuthMiddleware(jwtService))
	{
		protected.GET("/tasks", taskController.ListTasks)
		protected.GET("/tasks/:id", taskController.GetTask)
		protected.GET("/users/:id", userController.GetUserByID)
		admin := protected.Group("/")
		admin.Use(infrastructure.AdminMiddleware())
		{
			admin.POST("/tasks", taskController.CreateTask)
			admin.PUT("/tasks/:id", taskController.UpdateTask)
			admin.DELETE("/tasks/:id", taskController.DeleteTask)
			admin.PUT("/promote/:username", userController.PromoteUser)
		}
	}
	return r
}
