package database

import (
	"gorm.io/gorm"
)

// Paginate is a GORM scope for pagination
func Paginate(page, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if perPage < 1 || perPage > 100 {
			perPage = 20
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

// FilterByBusiness is a GORM scope for filtering by business ID
func FilterByBusiness(businessID interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("business_id = ?", businessID)
	}
}

// FilterByBranch is a GORM scope for filtering by branch ID
func FilterByBranch(branchID interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("branch_id = ?", branchID)
	}
}

// FilterActive is a GORM scope for filtering active records
func FilterActive() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true)
	}
}

// OrderByCreatedDesc is a GORM scope for ordering by created_at descending
func OrderByCreatedDesc() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}
}

// OrderByCreatedAsc is a GORM scope for ordering by created_at ascending
func OrderByCreatedAsc() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at ASC")
	}
}
