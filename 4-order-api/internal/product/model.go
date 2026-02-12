package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"uniqueIndex"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price"`
}

func NewProduct(name string, description string, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      pq.StringArray{},
		Price:       price,
	}
}
