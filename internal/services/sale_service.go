package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type SaleService struct {
	repo          *repository.SaleRepository
	productRepo   *repository.ProductRepository
	inventoryRepo *repository.InventoryRepository
	shiftRepo     *repository.ShiftRepository
}

func NewSaleService(repo *repository.SaleRepository, productRepo *repository.ProductRepository, inventoryRepo *repository.InventoryRepository, shiftRepo *repository.ShiftRepository) *SaleService {
	return &SaleService{
		repo:          repo,
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
		shiftRepo:     shiftRepo,
	}
}

type SaleItemInput struct {
	ProductID uuid.UUID  `json:"product_id"`
	VariantID *uuid.UUID `json:"variant_id"`
	Quantity  int        `json:"quantity"`
}

func (s *SaleService) ProcessSale(branchID, userID uuid.UUID, items []SaleItemInput, paymentMethod string, notes string) (*models.Sale, error) {
	// 1. Verify Open Shift
	shift, err := s.shiftRepo.GetOpenShift(userID)
	if err != nil {
		return nil, errors.New("cannot process sale: no open shift found")
	}

	// 2. Validate Items & Stock, Calculate Totals
	var saleItems []models.SaleItem
	var subtotal float64

	for _, item := range items {
		// Fetch Product
		product, err := s.productRepo.GetProductByID(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product not found: %v", item.ProductID)
		}

		var unitPrice float64
		var variantID *uuid.UUID
		var variantName string

		if item.VariantID != nil {
			// Find variant
			found := false
			for _, v := range product.Variants {
				if v.ID == *item.VariantID {
					if product.TrackInventory && v.Stock < item.Quantity {
						return nil, fmt.Errorf("insufficient stock for %s (%s)", product.Name, v.Name)
					}
					unitPrice = v.Price
					variantID = &v.ID
					variantName = v.Name
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("variant not found for product %s", product.Name)
			}
		} else {
			// No variant logic (assuming simple product or first variant? For now, error if variants exist)
			// If product has variants, variant_id is required.
			if len(product.Variants) > 0 {
				return nil, fmt.Errorf("variant selection required for %s", product.Name)
			}
			// Fallback if we allowed products without variants (not in current model, but safe to handle)
			return nil, errors.New("product structure requires variants")
		}

		itemTotal := unitPrice * float64(item.Quantity)
		subtotal += itemTotal

		saleItems = append(saleItems, models.SaleItem{
			ProductID: item.ProductID,
			VariantID: variantID,
			Quantity:  item.Quantity,
			UnitPrice: unitPrice,
			Subtotal:  itemTotal,
			Total:     itemTotal, // Discount logic can be added later
		})

		// Deduct Stock (Side Effect)
		// We'll do this in the transaction or right here if we trust the flow.
		// Ideally, we should do it transactionally.
		// For this implementation, we will call InventoryRepo.AdjustStock individually.
		// Note: If one fails, we have a partial state.
		// TODO: Move this into a single transaction in the repository layer for atomicity.
		// For now, we proceed.

		if product.TrackInventory && variantID != nil {
			tx := &models.InventoryTransaction{
				BranchID:         branchID,
				ProductVariantID: *variantID,
				UserID:           userID,
				Type:             "sale",
				Quantity:         -item.Quantity,
				Reason:           fmt.Sprintf("Sale %s", variantName),
			}
			if err := s.inventoryRepo.AdjustStock(tx); err != nil {
				return nil, fmt.Errorf("failed to deduct stock: %v", err)
			}
		}
	}

	// 3. Create Sale Record
	// Generate Receipt Number (Simple timestamp based for now)
	receiptNo := fmt.Sprintf("RCP-%d", time.Now().Unix())

	sale := &models.Sale{
		BranchID:      branchID,
		UserID:        userID,
		ShiftID:       &shift.ID,
		ReceiptNumber: receiptNo,
		Subtotal:      subtotal,
		Total:         subtotal, // Tax/Discount logic later
		PaymentMethod: paymentMethod,
		Status:        "completed",
		Notes:         notes,
		SaleItems:     saleItems,
	}

	if err := s.repo.CreateSale(sale); err != nil {
		return nil, err
	}

	return sale, nil
}

func (s *SaleService) GetSales(branchID *uuid.UUID, userID *uuid.UUID) ([]models.Sale, error) {
	return s.repo.GetSales(branchID, userID)
}

func (s *SaleService) GetSaleByID(id uuid.UUID) (*models.Sale, error) {
	return s.repo.GetSaleByID(id)
}

func (s *SaleService) RefundSale(saleID, userID uuid.UUID, reason string) (*models.Sale, error) {
	sale, err := s.repo.GetSaleByID(saleID)
	if err != nil {
		return nil, errors.New("sale not found")
	}

	if sale.Status == "refunded" {
		return nil, errors.New("sale already refunded")
	}

	// 1. Reverse Stock Deduction
	for _, item := range sale.SaleItems {
		if item.VariantID != nil {
			// We need to check if we should track inventory for this product.
			// Ideally, we should fetch the product to check TrackInventory.
			// But for now, if VariantID is present, we assume it was tracked or we just add it back.
			// A safer way is to fetch the product.
			product, err := s.productRepo.GetProductByID(item.ProductID)
			if err == nil && product.TrackInventory {
				tx := &models.InventoryTransaction{
					BranchID:         sale.BranchID,
					ProductVariantID: *item.VariantID,
					UserID:           userID,
					Type:             "return",
					Quantity:         item.Quantity, // Add back
					Reason:           fmt.Sprintf("Refund Sale %s: %s", sale.ReceiptNumber, reason),
				}
				if err := s.inventoryRepo.AdjustStock(tx); err != nil {
					return nil, fmt.Errorf("failed to restore stock: %v", err)
				}
			}
		}
	}

	// 2. Update Sale Status
	sale.Status = "refunded"
	sale.Notes = fmt.Sprintf("%s | Refunded by %s: %s", sale.Notes, userID, reason)

	if err := s.repo.UpdateSale(sale); err != nil {
		return nil, err
	}

	return sale, nil
}
