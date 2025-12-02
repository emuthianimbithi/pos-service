package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

// AdjustStock updates the stock level of a variant and logs the transaction
func (r *InventoryRepository) AdjustStock(transaction *models.InventoryTransaction) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Log the transaction
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// 2. Update the variant stock
		// Use gorm.Expr to handle concurrency safely (stock = stock + quantity)
		if err := tx.Model(&models.ProductVariant{}).
			Where("id = ?", transaction.ProductVariantID).
			Update("stock", gorm.Expr("stock + ?", transaction.Quantity)).
			Error; err != nil {
			return err
		}

		return nil
	})
}

// GetStockHistory retrieves transaction history for a variant
func (r *InventoryRepository) GetStockHistory(variantID uuid.UUID) ([]models.InventoryTransaction, error) {
	var history []models.InventoryTransaction
	err := r.db.Where("product_variant_id = ?", variantID).
		Order("created_at desc").
		Preload("User").
		Find(&history).Error
	return history, err
}
