package services

import (
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportService struct {
	db *gorm.DB
}

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{db: db}
}

func (s *ReportService) GetSalesReport(branchID *uuid.UUID, startDate, endDate time.Time) (*models.SalesReport, error) {
	query := s.db.Model(&models.Sale{})

	if branchID != nil {
		query = query.Where("branch_id = ?", *branchID)
	}
	query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)

	var totalSales, cashSales, cardSales, mpesaSales float64
	var totalOrders int64

	// Get summary
	query.Count(&totalOrders)
	query.Select("SUM(total) as total").Row().Scan(&totalSales)
	s.db.Model(&models.Sale{}).Where("payment_method = ?", "cash").Where("created_at BETWEEN ? AND ?", startDate, endDate).Select("SUM(total)").Row().Scan(&cashSales)
	s.db.Model(&models.Sale{}).Where("payment_method = ?", "card").Where("created_at BETWEEN ? AND ?", startDate, endDate).Select("SUM(total)").Row().Scan(&cardSales)
	s.db.Model(&models.Sale{}).Where("payment_method = ?", "mpesa").Where("created_at BETWEEN ? AND ?", startDate, endDate).Select("SUM(total)").Row().Scan(&mpesaSales)

	averageOrder := float64(0)
	if totalOrders > 0 {
		averageOrder = totalSales / float64(totalOrders)
	}

	return &models.SalesReport{
		StartDate:    startDate,
		EndDate:      endDate,
		TotalSales:   totalSales,
		TotalOrders:  totalOrders,
		CashSales:    cashSales,
		CardSales:    cardSales,
		MpesaSales:   mpesaSales,
		AverageOrder: averageOrder,
	}, nil
}

func (s *ReportService) GetInventoryReport(branchID uuid.UUID) (*models.InventoryReport, error) {
	var products []models.ProductInventoryReport

	rows, err := s.db.Table("product_variants").
		Select("products.id as product_id, products.name as product_name, product_variants.name as variant_name, product_variants.sku, product_variants.stock, product_variants.low_stock_threshold, product_variants.cost").
		Joins("JOIN products ON products.id = product_variants.product_id").
		Where("products.branch_id = ? AND products.is_active = ?", branchID, true).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalProducts := 0
	lowStockProducts := 0
	outOfStock := 0
	totalValue := float64(0)

	for rows.Next() {
		var p models.ProductInventoryReport
		var cost float64
		rows.Scan(&p.ProductID, &p.ProductName, &p.VariantName, &p.SKU, &p.Stock, &p.LowStockThreshold, &cost)

		p.Value = float64(p.Stock) * cost
		totalValue += p.Value

		if p.Stock == 0 {
			p.Status = "out_of_stock"
			outOfStock++
		} else if p.Stock <= p.LowStockThreshold {
			p.Status = "low_stock"
			lowStockProducts++
		} else {
			p.Status = "in_stock"
		}

		products = append(products, p)
		totalProducts++
	}

	return &models.InventoryReport{
		TotalProducts:    totalProducts,
		LowStockProducts: lowStockProducts,
		OutOfStock:       outOfStock,
		TotalValue:       totalValue,
		Products:         products,
	}, nil
}

func (s *ReportService) GetLowStockReport(branchID uuid.UUID) ([]models.ProductInventoryReport, error) {
	var products []models.ProductInventoryReport

	rows, err := s.db.Table("product_variants").
		Select("products.id as product_id, products.name as product_name, product_variants.name as variant_name, product_variants.sku, product_variants.stock, product_variants.low_stock_threshold, product_variants.cost").
		Joins("JOIN products ON products.id = product_variants.product_id").
		Where("products.branch_id = ? AND products.is_active = ? AND product_variants.stock <= product_variants.low_stock_threshold", branchID, true).
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.ProductInventoryReport
		var cost float64
		rows.Scan(&p.ProductID, &p.ProductName, &p.VariantName, &p.SKU, &p.Stock, &p.LowStockThreshold, &cost)

		p.Value = float64(p.Stock) * cost
		if p.Stock == 0 {
			p.Status = "out_of_stock"
		} else {
			p.Status = "low_stock"
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *ReportService) GetCustomerReport(branchID uuid.UUID, limit int) (*models.CustomerReport, error) {
	var totalCustomers int64
	s.db.Model(&models.Customer{}).Where("branch_id = ?", branchID).Count(&totalCustomers)

	var topCustomers []models.CustomerSalesReport
	err := s.db.Table("customers").
		Select("customers.id as customer_id, CONCAT(customers.first_name, ' ', customers.last_name) as customer_name, customers.total_spend as total_purchases, customers.loyalty_points").
		Where("customers.branch_id = ?", branchID).
		Order("customers.total_spend DESC").
		Limit(limit).
		Scan(&topCustomers).Error

	if err != nil {
		return nil, err
	}

	return &models.CustomerReport{
		TotalCustomers: int(totalCustomers),
		TopCustomers:   topCustomers,
	}, nil
}
