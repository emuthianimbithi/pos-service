package services

import (
	"errors"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type ShiftService struct {
	repo     *repository.ShiftRepository
	saleRepo *repository.SaleRepository
}

func NewShiftService(repo *repository.ShiftRepository, saleRepo *repository.SaleRepository) *ShiftService {
	return &ShiftService{
		repo:     repo,
		saleRepo: saleRepo,
	}
}

func (s *ShiftService) OpenShift(branchID, userID uuid.UUID, openingBalance float64) (*models.Shift, error) {
	// Check if user already has an open shift
	existing, err := s.repo.GetOpenShift(userID)
	if err == nil && existing != nil {
		return nil, errors.New("user already has an open shift")
	}

	shift := &models.Shift{
		BranchID:       branchID,
		UserID:         userID,
		StartTime:      time.Now(),
		OpeningBalance: openingBalance,
		Status:         "open",
	}

	if err := s.repo.CreateShift(shift); err != nil {
		return nil, err
	}

	return shift, nil
}

func (s *ShiftService) CloseShift(userID uuid.UUID, closingBalance float64, notes string) (*models.Shift, error) {
	shift, err := s.repo.GetOpenShift(userID)
	if err != nil {
		return nil, errors.New("no open shift found")
	}

	// Calculate sales totals
	now := time.Now()
	salesSummary, err := s.saleRepo.GetSalesSummary(userID, shift.StartTime, now)
	if err != nil {
		return nil, err
	}

	cashSales := salesSummary["cash"]
	mpesaSales := salesSummary["mpesa"]
	cardSales := salesSummary["card"]
	totalSales := cashSales + mpesaSales + cardSales

	// Calculate variance
	// Expected Cash = Opening Balance + Cash Sales
	expectedCash := shift.OpeningBalance + cashSales
	variance := closingBalance - expectedCash

	// Update shift
	shift.EndTime = &now
	shift.ClosingBalance = &closingBalance
	shift.CashSales = cashSales
	shift.MpesaSales = mpesaSales
	shift.CardSales = cardSales
	shift.TotalSales = totalSales
	shift.CashVariance = variance
	shift.Notes = notes
	shift.Status = "closed"

	if err := s.repo.UpdateShift(shift); err != nil {
		return nil, err
	}

	return shift, nil
}

func (s *ShiftService) GetShifts(branchID *uuid.UUID, userID *uuid.UUID) ([]models.Shift, error) {
	return s.repo.GetShifts(branchID, userID)
}
