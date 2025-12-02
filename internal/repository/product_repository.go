package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Categories
func (r *ProductRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *ProductRepository) GetCategories(branchID uuid.UUID) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("branch_id = ?", branchID).Find(&categories).Error
	return categories, err
}

// Products
func (r *ProductRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetProducts(branchID uuid.UUID) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("branch_id = ?", branchID).
		Preload("Category").
		Preload("Variants").
		Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").
		Preload("Variants").
		First(&product, "id = ?", id).Error
	return &product, err
}

// Variants
func (r *ProductRepository) CreateVariant(variant *models.ProductVariant) error {
	return r.db.Create(variant).Error
}

func (r *ProductRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Variants").Preload("Category").First(&product, "id = ?", id).Error
	return &product, err
}

func (r *ProductRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) GetVariantByID(id uuid.UUID) (*models.ProductVariant, error) {
	var variant models.ProductVariant
	err := r.db.Preload("Product").First(&variant, "id = ?", id).Error
	return &variant, err
}

func (r *ProductRepository) UpdateVariant(variant *models.ProductVariant) error {
	return r.db.Save(variant).Error
}

func (r *ProductRepository) DeleteVariant(id uuid.UUID) error {
	return r.db.Delete(&models.ProductVariant{}, "id = ?", id).Error
}

func (r *ProductRepository) GetVariantBySKU(sku string) (*models.ProductVariant, error) {
	var variant models.ProductVariant
	err := r.db.Where("sku = ?", sku).First(&variant).Error
	return &variant, err
}
