package repository

import (
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

// GetSalesSummary calculates total sales by payment method for a specific user and time range
func (r *SaleRepository) GetSalesSummary(userID uuid.UUID, startTime, endTime time.Time) (map[string]float64, error) {
	type Result struct {
		PaymentMethod string
		Total         float64
	}

	var results []Result
	err := r.db.Model(&models.Sale{}).
		Select("payment_method, SUM(total) as total").
		Where("user_id = ? AND created_at >= ? AND created_at <= ? AND status = ?", userID, startTime, endTime, "completed").
		Group("payment_method").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	summary := make(map[string]float64)
	for _, res := range results {
		summary[res.PaymentMethod] = res.Total
	}

	return summary, nil
}

// CreateSale creates a new sale and its items in a transaction
func (r *SaleRepository) CreateSale(sale *models.Sale) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Create Sale record
		if err := tx.Create(sale).Error; err != nil {
			return err
		}

		// 2. Create Sale Items (GORM handles this via association if set, but explicit is safer for errors)
		// Since SaleItems are in the struct, GORM's Create(sale) should handle it.
		// However, we might want to ensure stock deduction happens here if we move logic to repo,
		// but logic belongs in service. Service will call this.

		return nil
	})
}

// GetSales retrieves sales with filters
func (r *SaleRepository) GetSales(branchID *uuid.UUID, userID *uuid.UUID) ([]models.Sale, error) {
	query := r.db.Model(&models.Sale{}).
		Preload("User").
		Preload("SaleItems").
		Preload("SaleItems.Product").
		Order("created_at desc")

	if branchID != nil {
		query = query.Where("branch_id = ?", branchID)
	}
	if userID != nil {
		query = query.Where("user_id = ?", userID)
	}

	var sales []models.Sale
	err := query.Find(&sales).Error
	return sales, err
}

func (r *SaleRepository) GetSaleByID(id uuid.UUID) (*models.Sale, error) {
	var sale models.Sale
	err := r.db.Preload("User").
		Preload("SaleItems").
		Preload("SaleItems.Product").
		First(&sale, "id = ?", id).Error
	return &sale, err
}

func (r *SaleRepository) UpdateSale(sale *models.Sale) error {
	return r.db.Save(sale).Error
}
