package repository

import (
	"app/models"

	"gorm.io/gorm"
)

// UserRepository provides methods to interact with the User model in the database
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetAll fetches all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

// Create saves a new user to the database
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetByID fetches a user by ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// Update modifies an existing user's data
func (r *UserRepository) Update(id uint, user *models.User) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

// Delete removes a user from the database
func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}