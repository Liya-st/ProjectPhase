package usecases

import (
	"context"
	"errors"
	"log"
	"task_manager/domain"
	"task_manager/infrastructure"
)

type UserInput struct {
	Username string
	Password string
}

type UserUsecase struct {
	repo           domain.UserRepository
	passwordService infrastructure.PasswordService
	jwtService     infrastructure.JWTService
}

func NewUserUsecase(repo domain.UserRepository, ps infrastructure.PasswordService, js infrastructure.JWTService) *UserUsecase {
	return &UserUsecase{repo: repo, passwordService: ps, jwtService: js}
}

func (u *UserUsecase) RegisterUser(c context.Context, input UserInput) (*domain.User, error) {
	log.Printf("Registering user: %s", input.Username)
	_, err := u.repo.GetByUsername(c, input.Username)
	if err == nil {
		log.Printf("Username %s already exists", input.Username)
		return nil, domain.ErrUsernameExists
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		log.Printf("Error checking username: %v", err)
		return nil, err
	}

	hashedPassword, err := u.passwordService.HashPassword(input.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	count, err := u.repo.CountUsers(c)
	if err != nil {
		log.Printf("Error counting users: %v", err)
		return nil, err
	}
	role := domain.RoleUser
	if count == 0 {
		role = domain.RoleAdmin
		log.Printf("Assigning admin role to first user: %s", input.Username)
	}

	user, err := domain.NewUser(input.Username, hashedPassword, role)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	createdUser, err := u.repo.Create(c, user)
	if err != nil {
		log.Printf("Error saving user to database: %v", err)
		return nil, err
	}
	log.Printf("User %s registered successfully with role %s", createdUser.Username, createdUser.Role)
	return createdUser, nil
}

func (u *UserUsecase) LoginUser(c context.Context, username, password string) (string, error) {
	log.Printf("Attempting login for user: %s", username)
	user, err := u.repo.GetByUsername(c, username)
	if err != nil {
		log.Printf("User %s not found: %v", username, err)
		return "", domain.ErrUserNotFound
	}
	if !u.passwordService.ComparePassword(user.PasswordHash, password) {
		log.Printf("Invalid password for user %s", username)
		return "", domain.ErrInvalidPassword
	}

	token, err := u.jwtService.GenerateToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		log.Printf("Error generating token for user %s: %v", username, err)
		return "", err
	}
	log.Printf("User %s logged in successfully", username)
	return token, nil
}

func (u *UserUsecase) PromoteToAdmin(c context.Context, username, requesterID string) error {
	log.Printf("Promoting user %s by requester %s", username, requesterID)
	requester, err := u.repo.GetByID(c, requesterID)
	if err != nil {
		log.Printf("Requester %s not found: %v", requesterID, err)
		return domain.ErrUserNotFound
	}
	if requester.Role != domain.RoleAdmin {
		log.Printf("Requester %s is not admin", requesterID)
		return domain.ErrAdminRequired
	}

	user, err := u.repo.GetByUsername(c, username)
	if err != nil {
		log.Printf("User %s not found: %v", username, err)
		return domain.ErrUserNotFound
	}

	user.Role = domain.RoleAdmin
	err = u.repo.Update(c, user)
	if err != nil {
		log.Printf("Error updating user %s: %v", username, err)
		return err
	}
	log.Printf("User %s promoted to admin", username)
	return nil
}

func (u *UserUsecase) GetUserByID(c context.Context, id string) (*domain.User, error) {
	log.Printf("Fetching user by ID: %s", id)
	user, err := u.repo.GetByID(c, id)
	if err != nil {
		log.Printf("User %s not found: %v", id, err)
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}
