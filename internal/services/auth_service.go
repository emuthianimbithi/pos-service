package services

import (
	"errors"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   *repository.UserRepository
	config *config.Config
}

func NewAuthService(repo *repository.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		repo:   repo,
		config: cfg,
	}
}

// RegisterBusiness creates a new business and the initial admin user
func (s *AuthService) RegisterBusiness(name, email, password, firstName, lastName, phone string) (*models.User, error) {
	// Check if user exists
	if _, err := s.repo.GetUserByEmail(email); err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create Business
	business := &models.Business{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	if err := s.repo.CreateBusiness(business); err != nil {
		return nil, err
	}

	// Create Admin User
	user := &models.User{
		BusinessID:   business.ID,
		Email:        email,
		PasswordHash: string(hashedPassword),
		FirstName:    firstName,
		LastName:     lastName,
		Phone:        phone,
		Role:         "admin",
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(email, password string) (string, *models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", nil, errors.New("account is deactivated")
	}

	if !user.Business.IsActive {
		return "", nil, errors.New("business account is deactivated")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"business_id": user.BusinessID,
		"role":        user.Role,
		"branch_id":   user.BranchID,
		"exp":         time.Now().Add(time.Hour * time.Duration(s.config.JWTExpiration)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
