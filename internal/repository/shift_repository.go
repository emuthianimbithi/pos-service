package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) *ShiftRepository {
	return &ShiftRepository{db: db}
}

func (r *ShiftRepository) CreateShift(shift *models.Shift) error {
	return r.db.Create(shift).Error
}

func (r *ShiftRepository) UpdateShift(shift *models.Shift) error {
	return r.db.Save(shift).Error
}

func (r *ShiftRepository) GetOpenShift(userID uuid.UUID) (*models.Shift, error) {
	var shift models.Shift
	err := r.db.Where("user_id = ? AND status = ?", userID, "open").First(&shift).Error
	if err != nil {
		return nil, err
	}
	return &shift, nil
}

func (r *ShiftRepository) GetShiftByID(id uuid.UUID) (*models.Shift, error) {
	var shift models.Shift
	err := r.db.Preload("User").Preload("Branch").First(&shift, "id = ?", id).Error
	return &shift, err
}

// GetShifts retrieves shifts based on filters
func (r *ShiftRepository) GetShifts(branchID *uuid.UUID, userID *uuid.UUID) ([]models.Shift, error) {
	query := r.db.Model(&models.Shift{}).Preload("User").Order("created_at desc")

	if branchID != nil {
		query = query.Where("branch_id = ?", branchID)
	}
	if userID != nil {
		query = query.Where("user_id = ?", userID)
	}

	var shifts []models.Shift
	err := query.Find(&shifts).Error
	return shifts, err
}
