package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"index"`  // Adding index on Name
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"index"` // Adding index on Price
	Stock       int     `json:"stock"`
}

// TableName allows you to specify a custom table name for the User model
func (p *Product) TableName() string {
	return "products"
}
