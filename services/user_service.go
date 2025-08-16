package services

import (
	"app/models"
	"app/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(id uint, user *models.User) error {
	return s.Repo.Update(id, user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}