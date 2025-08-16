package services

import (
	"app/models"
	"app/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) GetAllProducts(page, pageSize int, filters map[string]interface{}) ([]models.Product, int, error) {
	// Directly delegate to the repository
	return s.Repo.GetAll(page, pageSize, filters)
}

func (s *ProductService) CreateProduct(p *models.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(id uint, p *models.Product) error {
	return s.Repo.Update(id, p)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.Delete(id)
}