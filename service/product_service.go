package service

import (
	"challenge-10/models"
	"challenge-10/repository"
)

type ProductService interface {
	GetProductByID(productID int) (*models.Product, error)
}

type ProductServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductServiceImpl(repo repository.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) GetProductByID(productID int) (*models.Product, error) {
	product, err := s.repo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductServiceImpl) GetAllProducts() ([]*models.Product, error) {
	products, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}