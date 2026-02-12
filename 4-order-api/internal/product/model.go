package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"uniqueIndex"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewProduct(name string, description string, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
}
