package constants

// User roles
const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleCashier = "cashier"
)

// Payment methods
const (
	PaymentCash   = "cash"
	PaymentCard   = "card"
	PaymentMPesa  = "mpesa"
	PaymentCredit = "credit"
)

// Sale statuses
const (
	SaleStatusPending   = "pending"
	SaleStatusCompleted = "completed"
	SaleStatusVoided    = "voided"
	SaleStatusRefunded  = "refunded"
)

// Pagination defaults
const (
	DefaultPage    = 1
	DefaultPerPage = 20
	MaxPerPage     = 100
)

// JWT defaults
const (
	DefaultJWTExpiration = 24 // hours
)
