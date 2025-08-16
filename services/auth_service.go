package services

import (
	"app/models"
	"app/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

// Register registers a new user after validating the data
func (s *AuthService) Register(email, password, firstName, lastName string) error {
	// Check if the user already exists
	existingUser, err := s.Repo.GetByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user model
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword), // Store the hashed password
		FirstName: firstName,
		LastName:  lastName,
	}

	// Save the user to the database
	return s.Repo.CreateUser(user)
}

// Login authenticates a user and generates a token (JWT)
func (s *AuthService) Login(email, password string) (string, error) {
	// Get the user by email
	user, err := s.Repo.GetByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a token (JWT) or a session token here
	// Placeholder: return a dummy token
	token := "JWT-TOKEN" // Replace this with actual JWT generation logic

	return token, nil
}
