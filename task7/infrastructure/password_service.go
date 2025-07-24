package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hash, password string) bool
}

type PasswordServiceImpl struct{}

func NewPasswordService() PasswordService {
	return &PasswordServiceImpl{}
}

func (s *PasswordServiceImpl) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (s *PasswordServiceImpl) ComparePassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
