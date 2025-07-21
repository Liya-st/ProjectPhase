package models

import (
	"time"
 "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
    AdminRole Role = "admin"
    UserRole  Role = "user"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username  string             `json:"username" validate:"required"`
    Email     string             `json:"email" validate:"required,email"`
    Password  string             `json:"password" validate:"required,min=6"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
    Token     string             `json:"token" bson:"token"`
    UserType  string             `json:"user_type" bson:"user_type" validate:"required,oneof=admin user"`
}


type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}