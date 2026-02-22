package product

import (
	"http/4-order-api/pkg/db"

	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{Database: database}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *ProductRepository) GetByName(name string) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
	var products []Product
	result := repo.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) Update(id uint, patch *Product) (*Product, error) {
	var updated Product

	result := repo.Database.DB.
		Model(&updated).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(patch)

	if result.Error != nil {
		return nil, result.Error
	}

	return &updated, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ProductRepository) GetByID(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, id) // WHERE id = ?
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
