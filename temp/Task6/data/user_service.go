package data

import (
	"context"
	"errors"
	"os"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
	secretKey  string
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		collection: db.Collection("users"),
		secretKey:  os.Getenv("JWT_SECRET"),
	}
}

func (s *UserService) RegisterUser(user *models.User) error {
	// Check if username exists
	var existingUser models.User
	err := s.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return errors.New("username already exists")
	}
	if err != mongo.ErrNoDocuments {
		return err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.ID = uuid.New().String()
	user.Password = string(hashedPassword)

	// Set first user as admin
	count, err := s.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	if count == 0 {
		user.Role = models.RoleAdmin
	} else {
		user.Role = models.RoleUser
	}

	_, err = s.collection.InsertOne(context.Background(), user)
	return err
}

func (s *UserService) Login(username, password string) (string, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) PromoteToAdmin(username string) error {
	result, err := s.collection.UpdateOne(
		context.Background(),
		bson.M{"username": username},
		bson.M{"$set": bson.M{"role": models.RoleAdmin}},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.User{}, errors.New("user not found")
	}
	return user, err
}