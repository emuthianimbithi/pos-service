package models

import (
	"time"

	"github.com/google/uuid"
)

// MpesaConfig stores M-Pesa credentials for a specific branch
type MpesaConfig struct {
	BranchModel // Includes ID, CreatedAt, UpdatedAt, DeletedAt, BusinessID, BranchID

	ConsumerKey    string `gorm:"type:varchar(255);not null" json:"consumer_key"`
	ConsumerSecret string `gorm:"type:varchar(255);not null" json:"consumer_secret"`
	PassKey        string `gorm:"type:varchar(255);not null" json:"pass_key"`
	ShortCode      string `gorm:"type:varchar(50);not null" json:"short_code"`
	Type           string `gorm:"type:varchar(50);default:'paybill'" json:"type"`        // paybill or till_number
	Environment    string `gorm:"type:varchar(50);default:'sandbox'" json:"environment"` // sandbox or production

	// Callback URLs (optional override, otherwise system defaults used)
	CallbackURL string `gorm:"type:varchar(255)" json:"callback_url"`
}

// MpesaTransaction tracks M-Pesa payments
type MpesaTransaction struct {
	BranchModel // Includes ID, CreatedAt, UpdatedAt, DeletedAt, BusinessID, BranchID

	SaleID *uuid.UUID `gorm:"type:uuid;index" json:"sale_id,omitempty"`

	// Request Details
	MerchantRequestID string  `gorm:"type:varchar(100);index" json:"merchant_request_id"`
	CheckoutRequestID string  `gorm:"type:varchar(100);index" json:"checkout_request_id"`
	Amount            float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	PhoneNumber       string  `gorm:"type:varchar(20);not null" json:"phone_number"`
	Reference         string  `gorm:"type:varchar(100)" json:"reference"`
	Description       string  `gorm:"type:varchar(255)" json:"description"`

	// Status
	Status          string    `gorm:"type:varchar(50);default:'pending'" json:"status"` // pending, completed, failed, cancelled
	ResultCode      int       `json:"result_code"`
	ResultDesc      string    `gorm:"type:text" json:"result_desc"`
	MpesaReceipt    string    `gorm:"type:varchar(50);index" json:"mpesa_receipt"`
	TransactionDate time.Time `json:"transaction_date"`
}

// TableName specifies the table name for MpesaConfig
func (MpesaConfig) TableName() string {
	return "mpesa_configs"
}

// TableName specifies the table name for MpesaTransaction
func (MpesaTransaction) TableName() string {
	return "mpesa_transactions"
}
