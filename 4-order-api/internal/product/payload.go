package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" gorm:"uniqueIndex" validate:"required"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" validate:"required"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"name" gorm:"uniqueIndex" validate:"required"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" validate:"required"`
}
