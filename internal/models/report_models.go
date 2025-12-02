package models

import (
	"time"

	"github.com/google/uuid"
)

// Report DTOs

type SalesReport struct {
	StartDate    time.Time            `json:"start_date"`
	EndDate      time.Time            `json:"end_date"`
	TotalSales   float64              `json:"total_sales"`
	TotalOrders  int64                `json:"total_orders"`
	CashSales    float64              `json:"cash_sales"`
	CardSales    float64              `json:"card_sales"`
	MpesaSales   float64              `json:"mpesa_sales"`
	AverageOrder float64              `json:"average_order"`
	TopProducts  []ProductSalesReport `json:"top_products"`
	SalesByDay   []DailySalesReport   `json:"sales_by_day"`
}

type ProductSalesReport struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Revenue     float64   `json:"revenue"`
}

type DailySalesReport struct {
	Date   string  `json:"date"`
	Sales  float64 `json:"sales"`
	Orders int64   `json:"orders"`
}

type InventoryReport struct {
	TotalProducts    int                      `json:"total_products"`
	LowStockProducts int                      `json:"low_stock_products"`
	OutOfStock       int                      `json:"out_of_stock"`
	TotalValue       float64                  `json:"total_value"`
	Products         []ProductInventoryReport `json:"products"`
}

type ProductInventoryReport struct {
	ProductID         uuid.UUID `json:"product_id"`
	ProductName       string    `json:"product_name"`
	VariantName       string    `json:"variant_name"`
	SKU               string    `json:"sku"`
	Stock             int       `json:"stock"`
	LowStockThreshold int       `json:"low_stock_threshold"`
	Status            string    `json:"status"` // in_stock, low_stock, out_of_stock
	Value             float64   `json:"value"`  // stock * cost
}

type RevenueReport struct {
	Period         string               `json:"period"`
	TotalRevenue   float64              `json:"total_revenue"`
	TotalCost      float64              `json:"total_cost"`
	GrossProfit    float64              `json:"gross_profit"`
	ProfitMargin   float64              `json:"profit_margin"` // percentage
	DailyBreakdown []DailyRevenueReport `json:"daily_breakdown"`
}

type DailyRevenueReport struct {
	Date         string  `json:"date"`
	Revenue      float64 `json:"revenue"`
	Cost         float64 `json:"cost"`
	Profit       float64 `json:"profit"`
	ProfitMargin float64 `json:"profit_margin"`
}

type CustomerReport struct {
	TotalCustomers int                   `json:"total_customers"`
	TopCustomers   []CustomerSalesReport `json:"top_customers"`
}

type CustomerSalesReport struct {
	CustomerID     uuid.UUID `json:"customer_id"`
	CustomerName   string    `json:"customer_name"`
	TotalPurchases float64   `json:"total_purchases"`
	OrderCount     int64     `json:"order_count"`
	LoyaltyPoints  int       `json:"loyalty_points"`
}
