package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Preload("Business").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateBusiness(business *models.Business) error {
	return r.db.Create(business).Error
}

func (r *UserRepository) CreateBranch(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Business").Preload("Branch").First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) List(businessID uuid.UUID, branchID *uuid.UUID, role string, offset, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{}).Where("business_id = ?", businessID)

	if branchID != nil {
		query = query.Where("branch_id = ?", *branchID)
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at desc").Offset(offset).Limit(limit).Preload("Branch").Find(&users).Error
	return users, total, err
}
