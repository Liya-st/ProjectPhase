package models

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID       string   `json:"id" bson:"_id"`
	Username string   `json:"username" bson:"username"`
	Password string   `json:"password" bson:"password"`
	Role     UserRole `json:"role" bson:"role"`
}