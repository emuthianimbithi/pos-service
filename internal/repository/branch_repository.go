package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BranchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{db: db}
}

func (r *BranchRepository) Create(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *BranchRepository) GetByID(id uuid.UUID) (*models.Branch, error) {
	var branch models.Branch
	err := r.db.Preload("Business").First(&branch, "id = ?", id).Error
	return &branch, err
}

func (r *BranchRepository) ListByBusiness(businessID uuid.UUID, offset, limit int) ([]models.Branch, int64, error) {
	var branches []models.Branch
	var total int64

	query := r.db.Model(&models.Branch{}).Where("business_id = ?", businessID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&branches).Error
	return branches, total, err
}

func (r *BranchRepository) Update(branch *models.Branch) error {
	return r.db.Save(branch).Error
}
