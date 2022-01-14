package service

import (
	"errors"
	"time"

	"github.com/justjundana/go-crud-mysql/helper"
	"github.com/justjundana/go-crud-mysql/models"
	"github.com/justjundana/go-crud-mysql/repository"
)

type ProductService interface {
	GetProductsService() ([]models.Product, error)
	GetProductService(id int) (models.Product, error)
	CreateProductService(input helper.CreateProductRequest) (models.Product, error)
	UpdateProductService(id int, input helper.EditProductRequest) (models.Product, error)
	DeleteProductService(id, currentUser int) (models.Product, error)
}

type productService struct {
	repository repository.ProductRepositoryInterface
}

func NewProductService(repository repository.ProductRepositoryInterface) *productService {
	return &productService{repository}
}

func (s *productService) GetProductsService() ([]models.Product, error) {
	products, err := s.repository.GetProducts()
	return products, err
}

func (s *productService) GetProductService(id int) (models.Product, error) {
	product, err := s.repository.GetProduct(id)
	return product, err
}

func (s *productService) CreateProductService(input helper.CreateProductRequest) (models.Product, error) {
	product := models.Product{}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.User.ID = input.User.ID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	createProduct, err := s.repository.CreateProduct(product)
	return createProduct, err
}

func (s *productService) UpdateProductService(id int, input helper.EditProductRequest) (models.Product, error) {
	product, err := s.repository.GetProduct(id)
	if err != nil {
		return product, err
	}

	if product.User.ID != input.User.ID {
		return product, errors.New("you don't have permission to update this product")
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.UpdatedAt = time.Now()

	updateProduct, err := s.repository.UpdateProduct(product)
	return updateProduct, err
}

func (s *productService) DeleteProductService(id, currentUser int) (models.Product, error) {
	productID, err := s.GetProductService(id)
	if err != nil {
		return productID, err
	}

	if productID.User.ID != currentUser {
		return productID, errors.New("you don't have permission to delete this product")
	}

	deleteProduct, err := s.repository.DeleteProduct(productID)
	return deleteProduct, err
}
