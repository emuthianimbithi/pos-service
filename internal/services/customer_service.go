package services

import (
	"errors"
	"strings"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(branchID uuid.UUID, firstName, lastName, phone, email string) (*models.Customer, error) {
	if firstName == "" || lastName == "" || phone == "" {
		return nil, errors.New("first name, last name, and phone are required")
	}

	// Check if customer with phone already exists in this branch
	existing, err := s.repo.GetByPhone(branchID, phone)
	if err == nil && existing != nil {
		return nil, errors.New("customer with this phone number already exists in this branch")
	}

	customer := &models.Customer{
		BranchID:  branchID,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Email:     email,
	}

	if err := s.repo.Create(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) GetCustomer(id uuid.UUID) (*models.Customer, error) {
	return s.repo.GetByID(id)
}

func (s *CustomerService) UpdateCustomer(id uuid.UUID, firstName, lastName, phone, email string) (*models.Customer, error) {
	customer, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if firstName != "" {
		customer.FirstName = firstName
	}
	if lastName != "" {
		customer.LastName = lastName
	}
	if phone != "" {
		// Check uniqueness if phone is changing
		if phone != customer.Phone {
			existing, err := s.repo.GetByPhone(customer.BranchID, phone)
			if err == nil && existing != nil {
				return nil, errors.New("customer with this phone number already exists in this branch")
			}
			customer.Phone = phone
		}
	}
	if email != "" {
		customer.Email = email
	}

	if err := s.repo.Update(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) DeleteCustomer(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *CustomerService) ListCustomers(branchID uuid.UUID, page, limit int, search string) ([]models.Customer, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(branchID, offset, limit, strings.TrimSpace(search))
}
