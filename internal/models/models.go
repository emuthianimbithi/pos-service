package models

// TODO: Define your domain models here
// Each model should include:
// - UUID primary key
// - Timestamps (CreatedAt, UpdatedAt, DeletedAt for soft deletes)
// - Proper GORM tags for database mapping
// - JSON tags for API responses
// - Indexes for frequently queried fields
// - Foreign key relationships

// Example models to implement:
//
// type Business struct {
// 	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Name        string `gorm:"type:varchar(255);not null" json:"name"`
// 	Email       string `gorm:"type:varchar(255);unique;not null" json:"email"`
// 	Phone       string `gorm:"type:varchar(20)" json:"phone"`
// 	Address     string `gorm:"type:text" json:"address"`
// 	TaxID       string `gorm:"type:varchar(50)" json:"tax_id"`
// 	LogoURL     string `gorm:"type:varchar(500)" json:"logo_url"`
// 	IsActive    bool   `gorm:"default:true" json:"is_active"`
//
// 	// Relationships
// 	Branches []Branch `gorm:"foreignKey:BusinessID" json:"branches,omitempty"`
// 	Users    []User   `gorm:"foreignKey:BusinessID" json:"users,omitempty"`
// }
//
// type Branch struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BusinessID uuid.UUID      `gorm:"type:uuid;not null;index" json:"business_id"`
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Name     string `gorm:"type:varchar(255);not null" json:"name"`
// 	Code     string `gorm:"type:varchar(50);unique;not null" json:"code"`
// 	Phone    string `gorm:"type:varchar(20)" json:"phone"`
// 	Address  string `gorm:"type:text" json:"address"`
// 	IsActive bool   `gorm:"default:true" json:"is_active"`
//
// 	// Relationships
// 	Business Business  `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
// 	Products []Product `gorm:"foreignKey:BranchID" json:"products,omitempty"`
// 	Sales    []Sale    `gorm:"foreignKey:BranchID" json:"sales,omitempty"`
// }
//
// type User struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BusinessID uuid.UUID      `gorm:"type:uuid;not null;index" json:"business_id"`
// 	BranchID   *uuid.UUID     `gorm:"type:uuid;index" json:"branch_id,omitempty"`
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Email        string `gorm:"type:varchar(255);unique;not null" json:"email"`
// 	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
// 	FirstName    string `gorm:"type:varchar(100);not null" json:"first_name"`
// 	LastName     string `gorm:"type:varchar(100);not null" json:"last_name"`
// 	Phone        string `gorm:"type:varchar(20)" json:"phone"`
// 	Role         string `gorm:"type:varchar(50);not null" json:"role"` // admin, manager, cashier
// 	IsActive     bool   `gorm:"default:true" json:"is_active"`
//
// 	// Relationships
// 	Business *Business `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
// 	Branch   *Branch   `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
// }
//
// type Category struct {
// 	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BranchID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"branch_id"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Name        string `gorm:"type:varchar(255);not null" json:"name"`
// 	Description string `gorm:"type:text" json:"description"`
//
// 	// Relationships
// 	Branch   Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
// 	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
// }
//
// type Product struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BranchID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"branch_id"`
// 	CategoryID *uuid.UUID     `gorm:"type:uuid;index" json:"category_id,omitempty"`
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
// 	SKU         string  `gorm:"type:varchar(100);unique;not null" json:"sku"`
// 	Description string  `gorm:"type:text" json:"description"`
// 	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
// 	Cost        float64 `gorm:"type:decimal(10,2)" json:"cost"`
// 	Stock       int     `gorm:"default:0" json:"stock"`
// 	MinStock    int     `gorm:"default:0" json:"min_stock"`
// 	ImageURL    string  `gorm:"type:varchar(500)" json:"image_url"`
// 	IsActive    bool    `gorm:"default:true" json:"is_active"`
//
// 	// Relationships
// 	Branch   Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
// 	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
// }
//
// type Customer struct {
// 	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BranchID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"branch_id"`
// 	CreatedAt time.Time      `json:"created_at"`
// 	UpdatedAt time.Time      `json:"updated_at"`
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	Name         string  `gorm:"type:varchar(255);not null" json:"name"`
// 	Email        string  `gorm:"type:varchar(255)" json:"email"`
// 	Phone        string  `gorm:"type:varchar(20)" json:"phone"`
// 	Address      string  `gorm:"type:text" json:"address"`
// 	LoyaltyPoints int    `gorm:"default:0" json:"loyalty_points"`
// 	CreditLimit  float64 `gorm:"type:decimal(10,2);default:0" json:"credit_limit"`
//
// 	// Relationships
// 	Branch Branch `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
// 	Sales  []Sale `gorm:"foreignKey:CustomerID" json:"sales,omitempty"`
// }
//
// type Sale struct {
// 	ID         uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
// 	BranchID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"branch_id"`
// 	CustomerID *uuid.UUID     `gorm:"type:uuid;index" json:"customer_id,omitempty"`
// 	UserID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"` // Cashier
// 	CreatedAt  time.Time      `json:"created_at"`
// 	UpdatedAt  time.Time      `json:"updated_at"`
// 	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
//
// 	ReceiptNumber string  `gorm:"type:varchar(100);unique;not null" json:"receipt_number"`
// 	Subtotal      float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
// 	Tax           float64 `gorm:"type:decimal(10,2);default:0" json:"tax"`
// 	Discount      float64 `gorm:"type:decimal(10,2);default:0" json:"discount"`
// 	Total         float64 `gorm:"type:decimal(10,2);not null" json:"total"`
// 	PaymentMethod string  `gorm:"type:varchar(50);not null" json:"payment_method"` // cash, card, mpesa
// 	Status        string  `gorm:"type:varchar(50);default:'completed'" json:"status"`
// 	Notes         string  `gorm:"type:text" json:"notes"`
//
// 	// Relationships
// 	Branch    Branch     `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
// 	Customer  *Customer  `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
// 	User      User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
// 	SaleItems []SaleItem `gorm:"foreignKey:SaleID" json:"sale_items,omitempty"`
// }
//
// type SaleItem struct {
// 	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
// 	SaleID    uuid.UUID `gorm:"type:uuid;not null;index" json:"sale_id"`
// 	ProductID uuid.UUID `gorm:"type:uuid;not null;index" json:"product_id"`
// 	CreatedAt time.Time `json:"created_at"`
//
// 	Quantity  int     `gorm:"not null" json:"quantity"`
// 	UnitPrice float64 `gorm:"type:decimal(10,2);not null" json:"unit_price"`
// 	Subtotal  float64 `gorm:"type:decimal(10,2);not null" json:"subtotal"`
// 	Discount  float64 `gorm:"type:decimal(10,2);default:0" json:"discount"`
// 	Total     float64 `gorm:"type:decimal(10,2);not null" json:"total"`
//
// 	// Relationships
// 	Sale    Sale    `gorm:"foreignKey:SaleID" json:"sale,omitempty"`
// 	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
// }
