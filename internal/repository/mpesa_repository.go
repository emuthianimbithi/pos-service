package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MpesaRepository struct {
	db *gorm.DB
}

func NewMpesaRepository(db *gorm.DB) *MpesaRepository {
	return &MpesaRepository{db: db}
}

// GetConfigByBranchID retrieves M-Pesa configuration for a branch
func (r *MpesaRepository) GetConfigByBranchID(branchID uuid.UUID) (*models.MpesaConfig, error) {
	var config models.MpesaConfig
	if err := r.db.Where("branch_id = ?", branchID).First(&config).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// UpsertConfig creates or updates M-Pesa configuration
func (r *MpesaRepository) UpsertConfig(config *models.MpesaConfig) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "branch_id"}},
		UpdateAll: true,
	}).Create(config).Error
}

// CreateTransaction creates a new M-Pesa transaction record
func (r *MpesaRepository) CreateTransaction(tx *models.MpesaTransaction) error {
	return r.db.Create(tx).Error
}

// GetTransactionByCheckoutRequestID retrieves a transaction by checkout request ID
func (r *MpesaRepository) GetTransactionByCheckoutRequestID(checkoutRequestID string) (*models.MpesaTransaction, error) {
	var tx models.MpesaTransaction
	if err := r.db.Where("checkout_request_id = ?", checkoutRequestID).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

// UpdateTransactionStatus updates the status of a transaction
func (r *MpesaRepository) UpdateTransactionStatus(id uuid.UUID, status string, resultCode int, resultDesc string) error {
	return r.db.Model(&models.MpesaTransaction{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":      status,
		"result_code": resultCode,
		"result_desc": resultDesc,
	}).Error
}

// GetTransactionByID retrieves a transaction by its ID
func (r *MpesaRepository) GetTransactionByID(id uuid.UUID) (*models.MpesaTransaction, error) {
	var tx models.MpesaTransaction
	if err := r.db.Where("id = ?", id).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}
