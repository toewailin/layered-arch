package repository

import (
	"app/models"

	"gorm.io/gorm"
)

// AuthRepository provides methods for authentication-related database queries
type AuthRepository struct {
	DB *gorm.DB
}

// NewAuthRepository creates a new AuthRepository instance
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

// GetByEmail fetches a user by their email
func (r *AuthRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser saves a new user to the database
func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
