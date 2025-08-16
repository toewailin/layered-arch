package services

import (
	"app/models"
	"app/repository"
)

type FaqService struct {
	Repo *repository.FaqRepository
}

func NewFaqService(repo *repository.FaqRepository) *FaqService {
	return &FaqService{Repo: repo}
}

func (s *FaqService) GetAll() ([]models.Faq, error) {
	return s.Repo.GetAll()
}

func (s *FaqService) Create(faq *models.Faq) error {
	return s.Repo.Create(faq)
}

func (s *FaqService) GetByID(id uint) (*models.Faq, error) {
	return s.Repo.GetByID(id)
}

func (s *FaqService) Update(id uint, faq *models.Faq) error {
	return s.Repo.Update(id, faq)
}

func (s *FaqService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
