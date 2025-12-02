package services

import (
	"errors"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(businessID uuid.UUID, branchID *uuid.UUID, role string, page, limit int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return s.repo.List(businessID, branchID, role, offset, limit)
}

func (s *UserService) CreateUser(businessID uuid.UUID, branchID *uuid.UUID, email, password, firstName, lastName, phone, role string) (*models.User, error) {
	// Check if user exists
	if _, err := s.repo.GetUserByEmail(email); err == nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		BusinessID:   businessID,
		Email:        email,
		PasswordHash: string(hashedPassword),
		FirstName:    firstName,
		LastName:     lastName,
		Phone:        phone,
		Role:         role,
		IsActive:     true,
	}

	if branchID != nil {
		user.BranchID = branchID
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserProfile(id uuid.UUID) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUserProfile(id uuid.UUID, firstName, lastName, phone string) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if phone != "" {
		user.Phone = phone
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// ChangePassword allows admin to reset password without requiring old password
func (s *UserService) ChangePassword(id uuid.UUID, newPassword string) error {
	if newPassword == "" || len(newPassword) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return s.repo.Update(user)
}

// DeactivateUser performs soft delete
func (s *UserService) DeactivateUser(id uuid.UUID) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	user.IsActive = false
	return s.repo.Update(user)
}
