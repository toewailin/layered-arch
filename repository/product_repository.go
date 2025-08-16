package repository

import (
	"app/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// GetAll with pagination and filtering
func (r *ProductRepository) GetAll(page, pageSize int, filters map[string]interface{}) ([]models.Product, int, error) {
	var products []models.Product
	var total int64

	// Start query building
	query := r.DB.Model(&models.Product{})

	// Apply filters dynamically
	for key, value := range filters {
		// Using exact match on the filter
		query = query.Where(key+" = ?", value)
	}

	// Get the total count for pagination
	query.Count(&total)

	// Paginate
	query = query.Offset((page - 1) * pageSize).Limit(pageSize)

	// Execute query
	err := query.Find(&products).Error
	return products, int(total), err
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Update(id uint, updated *models.Product) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updated).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}