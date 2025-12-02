package services

import (
	"errors"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type BranchService struct {
	repo *repository.BranchRepository
}

func NewBranchService(repo *repository.BranchRepository) *BranchService {
	return &BranchService{repo: repo}
}

func (s *BranchService) CreateBranch(businessID uuid.UUID, name, code, phone, address string) (*models.Branch, error) {
	if name == "" || code == "" {
		return nil, errors.New("name and code are required")
	}

	branch := &models.Branch{
		BusinessID: businessID,
		Name:       name,
		Code:       code,
		Phone:      phone,
		Address:    address,
		IsActive:   true,
	}

	if err := s.repo.Create(branch); err != nil {
		return nil, err
	}

	return branch, nil
}

func (s *BranchService) GetBranchByID(id, businessID uuid.UUID) (*models.Branch, error) {
	branch, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure branch belongs to the business
	if branch.BusinessID != businessID {
		return nil, errors.New("branch not found in this business")
	}

	return branch, nil
}

func (s *BranchService) ListBranches(businessID uuid.UUID, page, limit int) ([]models.Branch, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.ListByBusiness(businessID, offset, limit)
}

func (s *BranchService) UpdateBranch(id, businessID uuid.UUID, name, code, phone, address string) (*models.Branch, error) {
	branch, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure branch belongs to the business
	if branch.BusinessID != businessID {
		return nil, errors.New("branch not found in this business")
	}

	if name != "" {
		branch.Name = name
	}
	if code != "" {
		branch.Code = code
	}
	if phone != "" {
		branch.Phone = phone
	}
	if address != "" {
		branch.Address = address
	}

	if err := s.repo.Update(branch); err != nil {
		return nil, err
	}

	return branch, nil
}

// DeleteBranch performs soft delete by setting is_active to false
func (s *BranchService) DeleteBranch(id, businessID uuid.UUID) error {
	branch, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Ensure branch belongs to the business
	if branch.BusinessID != businessID {
		return errors.New("branch not found in this business")
	}

	branch.IsActive = false
	return s.repo.Update(branch)
}
