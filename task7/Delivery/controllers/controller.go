package controllers

import (
	"errors"
	"net/http"
	"task_manager/domain"
	"task_manager/usecases"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	usecase *usecases.TaskUsecase
}

type UserController struct {
	usecase *usecases.UserUsecase
}

func NewTaskController(usecase *usecases.TaskUsecase) *TaskController {
	return &TaskController{usecase: usecase}
}

func NewUserController(usecase *usecases.UserUsecase) *UserController {
	return &UserController{usecase: usecase}
}

func (controller *TaskController) CreateTask(c *gin.Context) {
	var input usecases.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	task, err := controller.usecase.CreateTask(c.Request.Context(), input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrInvalidTaskTitle) || errors.Is(err, domain.ErrInvalidTaskStatus) || errors.Is(err, domain.ErrInvalidDueDate) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (controller *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	task, err := controller.usecase.GetTask(c.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrInvalidTaskID) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (controller *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	var input usecases.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	task, err := controller.usecase.UpdateTask(c.Request.Context(), id, input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrInvalidTaskTitle) || errors.Is(err, domain.ErrInvalidTaskStatus) || errors.Is(err, domain.ErrInvalidDueDate) {
			status = http.StatusBadRequest
		} else if errors.Is(err, domain.ErrInvalidTaskID) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})
		return
	}

	err := controller.usecase.DeleteTask(c.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrInvalidTaskID) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (controller *TaskController) ListTasks(c *gin.Context) {
	tasks, err := controller.usecase.ListTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (controller *UserController) Register(c *gin.Context) {
	var input usecases.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	user, err := controller.usecase.RegisterUser(c.Request.Context(), input)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrUsernameExists) || errors.Is(err, domain.ErrInvalidUsername) || errors.Is(err, domain.ErrInvalidPassword) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}

func (controller *UserController) Login(c *gin.Context) {
	var input usecases.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	token, err := controller.usecase.LoginUser(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		status := http.StatusUnauthorized
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(status, gin.H{"error": "User not found"})
			return
		}
		if errors.Is(err, domain.ErrInvalidPassword) {
			c.JSON(status, gin.H{"error": "Invalid password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *UserController) PromoteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	requesterID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	err := controller.usecase.PromoteToAdmin(c.Request.Context(), username, requesterID.(string))
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrUserNotFound) || errors.Is(err, domain.ErrAdminRequired) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}

func (controller *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := controller.usecase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrUserNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	})
}
