package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"task_manager/data"
	"task_manager/models"
)

var validate = validator.New()
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))


// Register User

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := validate.Struct(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.CreateUser(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}


// Login User

func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	user, err := data.FindUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := generateJWT(user.ID.Hex(), user.Username, user.UserType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID.Hex(),
			"username": user.Username,
			"email":    user.Email,
			"userType": user.UserType,
		},
	})
}


// Get Profile

func GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(primitive.ObjectID)

	user, err := data.FindUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID.Hex(),
		"username":  user.Username,
		"email":     user.Email,
		"user_type": user.UserType,
		"created":   user.CreatedAt,
	})
}


func generateJWT(userID string, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role, // critical for AdminMiddleware
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}



// Create Task

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	// userID := c.MustGet("userID").(primitive.ObjectID)
	// task.UserID = userID
	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := validate.Struct(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}


// Get All Tasks for User

func GetTasks(c *gin.Context) {
	// userID := c.MustGet("userID").(primitive.ObjectID)

	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}


// Get Single Task by ID

func GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	// userID := c.MustGet("userID").(primitive.ObjectID)

	task, err := data.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}


// Update Task

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}

	if err := validate.StructPartial(&task, "Title", "Status", "DueDate"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID := c.MustGet("userID").(primitive.ObjectID)
	if err := data.UpdateTask(taskID, &task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not owned by user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}


// Delete Task

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	// userID := c.MustGet("userID").(primitive.ObjectID)

	if err := data.DeleteTask(taskID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or not owned by user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
