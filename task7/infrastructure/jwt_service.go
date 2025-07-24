package infrastructure

import (
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type JWTServiceImpl struct {
	secret string
}

func NewJWTService(secret string) JWTService {
	return &JWTServiceImpl{secret: secret}
}

func (s *JWTServiceImpl) GenerateToken(id, username, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *JWTServiceImpl) ValidateToken(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
