package repository

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) GetByID(id uuid.UUID) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Preload("Branch").First(&customer, "id = ?", id).Error
	return &customer, err
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Customer{}, "id = ?", id).Error
}

func (r *CustomerRepository) List(branchID uuid.UUID, offset, limit int, search string) ([]models.Customer, int64, error) {
	var customers []models.Customer
	var total int64

	query := r.db.Model(&models.Customer{}).Where("branch_id = ?", branchID)

	if search != "" {
		search = "%" + search + "%"
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR phone ILIKE ? OR email ILIKE ?", search, search, search, search)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&customers).Error
	return customers, total, err
}

func (r *CustomerRepository) GetByPhone(branchID uuid.UUID, phone string) (*models.Customer, error) {
	var customer models.Customer
	err := r.db.Where("branch_id = ? AND phone = ?", branchID, phone).First(&customer).Error
	return &customer, err
}
