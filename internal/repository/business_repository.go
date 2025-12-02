package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) *BusinessRepository {
	return &BusinessRepository{db: db}
}

func (r *BusinessRepository) Create(business *models.Business) error {
	return r.db.Create(business).Error
}

func (r *BusinessRepository) GetByID(id uuid.UUID) (*models.Business, error) {
	var business models.Business
	err := r.db.Preload("Branches").First(&business, "id = ?", id).Error
	return &business, err
}

func (r *BusinessRepository) List(offset, limit int) ([]models.Business, int64, error) {
	var businesses []models.Business
	var total int64

	query := r.db.Model(&models.Business{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&businesses).Error
	return businesses, total, err
}

func (r *BusinessRepository) Update(business *models.Business) error {
	return r.db.Save(business).Error
}
