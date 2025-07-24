package infrastructure

import (
	"log"
	"net/http"
	"strings"
	"task_manager/domain"

	"github.com/gin-gonic/gin"
)

type JWTService interface {
	GenerateToken(id, username, role string) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}

func AuthMiddleware(jwtService JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Println("Authorization header format must be Bearer {token}")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		token := parts[1]

		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			log.Printf("Invalid token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}


		userID, ok := claims["id"].(string)
		if !ok {
			log.Println("Invalid token claims: missing or invalid id")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			log.Println("Invalid token claims: missing or invalid role")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		log.Printf("Authenticated userID: %s, role: %s", userID, role)
		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		log.Printf("Checking admin access, role: %v, exists: %v", role, exists)
		if !exists || role != string(domain.RoleAdmin) {
			log.Println("Admin access denied")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			return
		}
		log.Println("Admin access granted")
		c.Next()
	}
}
