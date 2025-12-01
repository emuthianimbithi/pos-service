package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrInvalidUUID       = errors.New("invalid UUID format")
	ErrMissingUserID     = errors.New("user ID not found in context")
	ErrMissingBusinessID = errors.New("business ID not found in context")
	ErrMissingBranchID   = errors.New("branch ID not found in context")
)

// GetUserID extracts user ID from context
func GetUserID(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, ErrMissingUserID
	}

	uid, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrInvalidUUID
	}

	return uid, nil
}

// GetBusinessID extracts business ID from context
func GetBusinessID(c *gin.Context) (uuid.UUID, error) {
	businessID, exists := c.Get("business_id")
	if !exists {
		return uuid.Nil, ErrMissingBusinessID
	}

	bid, ok := businessID.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrInvalidUUID
	}

	return bid, nil
}

// GetBranchID extracts branch ID from context
func GetBranchID(c *gin.Context) (uuid.UUID, error) {
	branchID, exists := c.Get("branch_id")
	if !exists {
		return uuid.Nil, ErrMissingBranchID
	}

	bid, ok := branchID.(uuid.UUID)
	if !ok {
		return uuid.Nil, ErrInvalidUUID
	}

	return bid, nil
}

// GetRole extracts user role from context
func GetRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		return ""
	}

	r, ok := role.(string)
	if !ok {
		return ""
	}

	return r
}

// ParseUUIDParam parses UUID from URL parameter
func ParseUUIDParam(c *gin.Context, param string) (uuid.UUID, error) {
	idStr := c.Param(param)
	if idStr == "" {
		return uuid.Nil, errors.New("missing " + param + " parameter")
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, ErrInvalidUUID
	}

	return id, nil
}
