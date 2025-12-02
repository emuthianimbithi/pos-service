package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	RoleAdmin       = "admin"
	RoleManager     = "manager"
	RoleCashier     = "cashier"
	RoleStockKeeper = "stock_keeper"
)

type Business struct {
	BaseModel
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone    string `gorm:"type:varchar(20)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	TaxID    string `gorm:"type:varchar(50)" json:"tax_id"`
	LogoURL  string `gorm:"type:varchar(500)" json:"logo_url"`
	IsActive bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Branches []Branch `gorm:"foreignKey:BusinessID" json:"branches,omitempty"`
	Users    []User   `gorm:"foreignKey:BusinessID" json:"users,omitempty"`
}

type Category struct {
	BaseModel
	BranchID    uuid.UUID `gorm:"type:uuid;not null;index" json:"branch_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`

	// Relationships
	Branch   Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

type Product struct {
	BaseModel
	BranchID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"branch_id"`
	CategoryID *uuid.UUID `gorm:"type:uuid;index" json:"category_id,omitempty"`

	Name           string `gorm:"type:varchar(255);not null" json:"name"`
	Description    string `gorm:"type:text" json:"description"`
	ImageURL       string `gorm:"type:varchar(500)" json:"image_url"`
	IsActive       bool   `gorm:"default:true" json:"is_active"`
	TrackInventory bool   `gorm:"default:true" json:"track_inventory"`

	// Relationships
	Branch   Branch           `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Category *Category        `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Variants []ProductVariant `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
}

type ProductVariant struct {
	BaseModel
	ProductID uuid.UUID `gorm:"type:uuid;not null;index" json:"product_id"`

	Name              string  `gorm:"type:varchar(255);not null" json:"name"` // e.g., "Small Red"
	SKU               string  `gorm:"type:varchar(100);unique;not null" json:"sku"`
	Price             float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Cost              float64 `gorm:"type:decimal(10,2)" json:"cost"`
	Stock             int     `gorm:"default:0" json:"stock"`
	LowStockThreshold int     `gorm:"default:10" json:"low_stock_threshold"`

	// Relationships
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

type InventoryTransaction struct {
	BaseModel
	BranchID         uuid.UUID `gorm:"type:uuid;not null;index" json:"branch_id"`
	ProductVariantID uuid.UUID `gorm:"type:uuid;not null;index" json:"product_variant_id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"` // Who made the change

	Type        string     `gorm:"type:varchar(50);not null" json:"type"` // adjustment, sale, return, transfer_in, transfer_out
	Quantity    int        `gorm:"not null" json:"quantity"`              // Positive or negative
	Reason      string     `gorm:"type:varchar(255)" json:"reason"`
	ReferenceID *uuid.UUID `gorm:"type:uuid;index" json:"reference_id,omitempty"` // SaleID or TransferID

	// Relationships
	Branch         Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	ProductVariant ProductVariant `gorm:"foreignKey:ProductVariantID" json:"variant,omitempty"`
	User           User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type Branch struct {
	BaseModel
	BusinessID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_branch_code_business" json:"business_id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Code       string    `gorm:"type:varchar(50);not null;uniqueIndex:idx_branch_code_business" json:"code"`
	Phone      string    `gorm:"type:varchar(20)" json:"phone"`
	Address    string    `gorm:"type:text" json:"address"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`

	// Relationships
	Business Business `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
}

type User struct {
	BaseModel
	BusinessID uuid.UUID  `gorm:"type:uuid;not null;index:idx_user_business_role;index:idx_user_business_branch" json:"business_id"`
	BranchID   *uuid.UUID `gorm:"type:uuid;index:idx_user_business_branch" json:"branch_id,omitempty"`

	Email        string `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
	FirstName    string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string `gorm:"type:varchar(100);not null" json:"last_name"`
	Phone        string `gorm:"type:varchar(20)" json:"phone"`
	Role         string `gorm:"type:varchar(50);not null;index:idx_user_business_role" json:"role"` // admin, manager, cashier, stock_keeper
	IsActive     bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Business *Business `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
	Branch   *Branch   `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}

type Shift struct {
	BaseModel
	BranchID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"branch_id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	StartTime time.Time  `gorm:"not null" json:"start_time"`
	EndTime   *time.Time `json:"end_time,omitempty"`

	OpeningBalance float64  `gorm:"type:decimal(10,2);not null" json:"opening_balance"`
	ClosingBalance *float64 `gorm:"type:decimal(10,2)" json:"closing_balance,omitempty"`

	// Calculated fields (filled when closing)
	CashSales  float64 `gorm:"type:decimal(10,2);default:0" json:"cash_sales"`
	MpesaSales float64 `gorm:"type:decimal(10,2);default:0" json:"mpesa_sales"`
	CardSales  float64 `gorm:"type:decimal(10,2);default:0" json:"card_sales"`
	TotalSales float64 `gorm:"type:decimal(10,2);default:0" json:"total_sales"`

	CashVariance float64 `gorm:"type:decimal(10,2);default:0" json:"cash_variance"` // Actual Closing - (Opening + CashSales)
	Notes        string  `gorm:"type:text" json:"notes"`
	Status       string  `gorm:"type:varchar(20);default:'open'" json:"status"` // open, closed

	// Relationships
	Branch Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type Customer struct {
	BaseModel
	BranchID  uuid.UUID `gorm:"type:uuid;not null;index" json:"branch_id"`
	FirstName string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Phone     string    `gorm:"type:varchar(20);not null" json:"phone"`

	LoyaltyPoints int     `gorm:"default:0" json:"loyalty_points"`
	TotalSpend    float64 `gorm:"type:decimal(10,2);default:0" json:"total_spend"`

	// Relationships
	Branch Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
}

type Sale struct {
	BaseModel
	BranchID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"branch_id"`
	CustomerID *uuid.UUID `gorm:"type:uuid;index" json:"customer_id,omitempty"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`   // Cashier
	ShiftID    *uuid.UUID `gorm:"type:uuid;index" json:"shift_id,omitempty"` // Link sale to a shift

	ReceiptNumber string  `gorm:"type:varchar(100);unique;not null" json:"receipt_number"`
	Subtotal      float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	Tax           float64 `gorm:"type:decimal(10,2);default:0" json:"tax"`
	Discount      float64 `gorm:"type:decimal(10,2);default:0" json:"discount"`
	Total         float64 `gorm:"type:decimal(10,2);not null" json:"total"`
	PaymentMethod string  `gorm:"type:varchar(50);not null" json:"payment_method"` // cash, card, mpesa
	Status        string  `gorm:"type:varchar(50);default:'completed'" json:"status"`
	Notes         string  `gorm:"type:text" json:"notes"`

	// Relationships
	Branch    Branch     `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Customer  *Customer  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	SaleItems []SaleItem `gorm:"foreignKey:SaleID" json:"sale_items,omitempty"`
}

type SaleItem struct {
	BaseModel
	SaleID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"sale_id"`
	ProductID uuid.UUID  `gorm:"type:uuid;not null;index" json:"product_id"`
	VariantID *uuid.UUID `gorm:"type:uuid;index" json:"variant_id,omitempty"`

	Quantity  int     `gorm:"not null" json:"quantity"`
	UnitPrice float64 `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	Subtotal  float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	Discount  float64 `gorm:"type:decimal(10,2);default:0" json:"discount"`
	Total     float64 `gorm:"type:decimal(10,2);not null" json:"total"`

	// Relationships
	Sale    Sale    `gorm:"foreignKey:SaleID" json:"sale,omitempty"`
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
