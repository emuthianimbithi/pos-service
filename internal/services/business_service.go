package services

import (
	"errors"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type BusinessService struct {
	repo *repository.BusinessRepository
}

func NewBusinessService(repo *repository.BusinessRepository) *BusinessService {
	return &BusinessService{repo: repo}
}

func (s *BusinessService) CreateBusiness(name, email, phone, address, taxID string) (*models.Business, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	business := &models.Business{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Address:  address,
		TaxID:    taxID,
		IsActive: true,
	}

	if err := s.repo.Create(business); err != nil {
		return nil, err
	}

	return business, nil
}

func (s *BusinessService) GetBusinessByID(id uuid.UUID) (*models.Business, error) {
	return s.repo.GetByID(id)
}

func (s *BusinessService) ListBusinesses(page, limit int) ([]models.Business, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(offset, limit)
}

func (s *BusinessService) UpdateBusiness(id uuid.UUID, name, email, phone, address, taxID, logoURL string) (*models.Business, error) {
	business, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		business.Name = name
	}
	if email != "" {
		business.Email = email
	}
	if phone != "" {
		business.Phone = phone
	}
	if address != "" {
		business.Address = address
	}
	if taxID != "" {
		business.TaxID = taxID
	}
	if logoURL != "" {
		business.LogoURL = logoURL
	}

	if err := s.repo.Update(business); err != nil {
		return nil, err
	}

	return business, nil
}

// DeleteBusiness performs soft delete by setting is_active to false
func (s *BusinessService) DeleteBusiness(id uuid.UUID) error {
	business, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	business.IsActive = false
	return s.repo.Update(business)
}
