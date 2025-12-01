package models

// This file shows example usage of the base models
// Delete this file when you start implementing your actual models

/*
Example 1: Business model (top-level tenant)
Uses BaseModel since it's not scoped to another tenant

type Business struct {
	BaseModel  // Includes: ID, CreatedAt, UpdatedAt, DeletedAt

	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone    string `gorm:"type:varchar(20)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	IsActive bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Branches []Branch `gorm:"foreignKey:BusinessID" json:"branches,omitempty"`
}

Example 2: Branch model (scoped to business)
Uses TenantModel which includes BusinessID

type Branch struct {
	TenantModel  // Includes: BaseModel + BusinessID

	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Code     string `gorm:"type:varchar(50);unique;not null" json:"code"`
	Phone    string `gorm:"type:varchar(20)" json:"phone"`
	Address  string `gorm:"type:text" json:"address"`
	IsActive bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Business Business  `gorm:"foreignKey:BusinessID" json:"business,omitempty"`
	Products []Product `gorm:"foreignKey:BranchID" json:"products,omitempty"`
}

Example 3: Product model (scoped to branch)
Uses BranchModel which includes BusinessID and BranchID

type Product struct {
	BranchModel  // Includes: BaseModel + BusinessID + BranchID

	CategoryID  *uuid.UUID `gorm:"type:uuid;index" json:"category_id,omitempty"`
	Name        string     `gorm:"type:varchar(255);not null" json:"name"`
	SKU         string     `gorm:"type:varchar(100);unique;not null" json:"sku"`
	Description string     `gorm:"type:text" json:"description"`
	Price       float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	Cost        float64    `gorm:"type:decimal(10,2)" json:"cost"`
	Stock       int        `gorm:"default:0" json:"stock"`
	MinStock    int        `gorm:"default:0" json:"min_stock"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`

	// Relationships
	Branch   Branch    `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

Benefits of using base models:
1. Consistent UUID primary keys across all tables
2. Automatic timestamp management (CreatedAt, UpdatedAt)
3. Soft delete support (DeletedAt)
4. Automatic UUID generation via BeforeCreate hook
5. Tenant isolation built-in (BusinessID, BranchID)
6. Less boilerplate code
7. Easier to maintain and update common fields
*/
