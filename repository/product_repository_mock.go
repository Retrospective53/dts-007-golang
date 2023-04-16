package repository

import (
	"challenge-10/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProductByID(productID int) (*models.Product, error) {
	args := m.Called(productID)
	product := args.Get(0)
	err := args.Error(1)
	if product == nil {
		return nil, err
	}
	return product.(*models.Product), err
}

func (m *MockProductRepository) GetAllProducts() ([]*models.Product, error) {
	args := m.Called()
	products := args.Get(0)
	err := args.Error(1)
	if products == nil {
		return nil, err
	}
	return products.([]*models.Product), err
}

// func TestProductService_GetProductByID_Found(t *testing.T) {
// 	mockRepo := new(MockProductRepository)
// 	product := &models.Product{
// 		Title:       "Test Product",
// 		Description: "A test product",
// 		UserID:      1,
// 		User: &models.User{
// 			Fullname:  "Test User",
// 			Email: "testuser@example.com",
// 		},
// 	}
// 	mockRepo.On("GetProductByID", 1).Return(product, nil)
// 	service := NewProductServiceImpl(mockRepo)

// 	result, err := service.GetProductByID(1)

// 	assert.Nil(t, err)
// 	assert.Equal(t, product, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestProductService_GetProductByID_NotFound(t *testing.T) {
// 	mockRepo := new(MockProductRepository)
// 	mockRepo.On("GetProductByID", 1).Return(nil, errors.New("product not found"))
// 	service := NewProductServiceImpl(mockRepo)

// 	result, err := service.GetProductByID(1)

// 	assert.NotNil(t, err)
// 	assert.Nil(t, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestProductService_GetAllProducts_Found(t *testing.T) {
// 	mockRepo := new(MockProductRepository)
// 	products := []*models.Product{
// 		{
// 			Title:       "Test Product 1",
// 			Description: "A test product",
// 			UserID:      1,
// 			User: &models.User{
// 				Fullname:  "Test User 1",
// 				Email: "testuser1@example.com",
// 			},
// 		},
// 		{
// 			Title:       "Test Product 2",
// 			Description: "Another test product",
// 			UserID:      2,
// 			User: &models.User{
// 				Fullname:  "Test User 2",
// 				Email: "testuser2@example.com",
// 			},
// 		},
// 	}
// 	mockRepo.On("GetAllProducts").Return(products, nil)
// 	service := NewProductServiceImpl(mockRepo)

// 	result, err := service.GetAllProducts()

// 	assert.Nil(t, err)
// 	assert.Equal(t, products, result)
// 	mockRepo.AssertExpectations(t)
// }

// func TestProductService_GetAllProducts_NotFound(t *testing.T) {
// 	mockRepo := new(MockProductRepository)
// 	mockRepo.On("GetAllProducts").Return(nil, errors.New("no products found"))
// 	service := NewProductServiceImpl(mockRepo)

// 	result, err := service.GetAllProducts()

// 	assert.NotNil(t, err)
// 	assert.Nil(t, result)
// 	mockRepo.AssertExpectations(t)
// }