package repository

import (
	"challenge-10/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(productID int) (*models.Product, error)
	GetAllProducts() ([]*models.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetProductByID(productID int) (*models.Product, error) {
	product := &models.Product{}
	err := r.db.First(product, "id = ?", productID).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}