package repository

import (
	"app/models"

	"gorm.io/gorm"
)

type FaqRepository struct {
	DB *gorm.DB
}

func NewFaqRepository(db *gorm.DB) *FaqRepository {
	return &FaqRepository{DB: db}
}

func (r *FaqRepository) GetAll() ([]models.Faq, error) {
	var faqs []models.Faq
	err := r.DB.Find(&faqs).Error
	return faqs, err
}

func (r *FaqRepository) Create(faq *models.Faq) error {
	return r.DB.Create(faq).Error
}

func (r *FaqRepository) GetByID(id uint) (*models.Faq, error) {
	var faq models.Faq
	err := r.DB.First(&faq, id).Error
	return &faq, err
}

func (r *FaqRepository) Update(id uint, updated *models.Faq) error {
	return r.DB.Model(&models.Faq{}).Where("id = ?", id).Updates(updated).Error
}

func (r *FaqRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Faq{}, id).Error
}
