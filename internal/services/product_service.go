package services

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Categories
func (s *ProductService) CreateCategory(branchID uuid.UUID, name, description string) (*models.Category, error) {
	category := &models.Category{
		BranchID:    branchID,
		Name:        name,
		Description: description,
	}
	if err := s.repo.CreateCategory(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *ProductService) GetCategories(branchID uuid.UUID) ([]models.Category, error) {
	return s.repo.GetCategories(branchID)
}

// Products
func (s *ProductService) CreateProduct(branchID uuid.UUID, categoryID *uuid.UUID, name, description, imageURL string, trackInventory bool) (*models.Product, error) {
	product := &models.Product{
		BranchID:       branchID,
		CategoryID:     categoryID,
		Name:           name,
		Description:    description,
		ImageURL:       imageURL,
		TrackInventory: trackInventory,
	}
	if err := s.repo.CreateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProducts(branchID uuid.UUID) ([]models.Product, error) {
	return s.repo.GetProducts(branchID)
}

// Variants
func (s *ProductService) AddVariant(productID uuid.UUID, name, sku string, price, cost float64, stock, lowStockThreshold int) (*models.ProductVariant, error) {
	variant := &models.ProductVariant{
		ProductID:         productID,
		Name:              name,
		SKU:               sku,
		Price:             price,
		Cost:              cost,
		Stock:             stock,
		LowStockThreshold: lowStockThreshold,
	}
	if err := s.repo.CreateVariant(variant); err != nil {
		return nil, err
	}
	return variant, nil
}

func (s *ProductService) GetProduct(id uuid.UUID) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) UpdateProduct(id uuid.UUID, categoryID *uuid.UUID, name, description, imageURL string, trackInventory bool) (*models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if imageURL != "" {
		product.ImageURL = imageURL
	}
	if categoryID != nil {
		product.CategoryID = categoryID
	}
	product.TrackInventory = trackInventory

	if err := s.repo.UpdateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id uuid.UUID) error {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	product.IsActive = false
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) UpdateVariant(id uuid.UUID, name, sku string, price, cost float64, stock, lowStockThreshold int) (*models.ProductVariant, error) {
	variant, err := s.repo.GetVariantByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		variant.Name = name
	}
	if sku != "" {
		variant.SKU = sku
	}
	if price > 0 {
		variant.Price = price
	}
	if cost >= 0 {
		variant.Cost = cost
	}
	if stock >= 0 {
		variant.Stock = stock
	}
	if lowStockThreshold >= 0 {
		variant.LowStockThreshold = lowStockThreshold
	}

	if err := s.repo.UpdateVariant(variant); err != nil {
		return nil, err
	}
	return variant, nil
}

func (s *ProductService) DeleteVariant(id uuid.UUID) error {
	return s.repo.DeleteVariant(id)
}
