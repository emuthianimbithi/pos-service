package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface
func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return ""
	}
	var messages []string
	for _, err := range v {
		messages = append(messages, err.Field+": "+err.Message)
	}
	return strings.Join(messages, "; ")
}

// ToMap converts validation errors to a map
func (v ValidationErrors) ToMap() map[string]string {
	result := make(map[string]string)
	for _, err := range v {
		result[err.Field] = err.Message
	}
	return result
}

// Validator helps build validation rules
type Validator struct {
	errors ValidationErrors
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		errors: make(ValidationErrors, 0),
	}
}

// Required checks if a string field is not empty
func (v *Validator) Required(field, value string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " is required",
		})
	}
	return v
}

// Email validates email format
func (v *Validator) Email(field, value string) *Validator {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if value != "" && !emailRegex.MatchString(value) {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " must be a valid email address",
		})
	}
	return v
}

// MinLength checks minimum string length
func (v *Validator) MinLength(field, value string, min int) *Validator {
	if len(value) < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " must be at least " + string(rune(min)) + " characters",
		})
	}
	return v
}

// MaxLength checks maximum string length
func (v *Validator) MaxLength(field, value string, max int) *Validator {
	if len(value) > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " must be at most " + string(rune(max)) + " characters",
		})
	}
	return v
}

// Min checks minimum numeric value
func (v *Validator) Min(field string, value, min float64) *Validator {
	if value < min {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " must be at least " + string(rune(int(min))),
		})
	}
	return v
}

// Max checks maximum numeric value
func (v *Validator) Max(field string, value, max float64) *Validator {
	if value > max {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: field + " must be at most " + string(rune(int(max))),
		})
	}
	return v
}

// UUID validates UUID format
func (v *Validator) UUID(field, value string) *Validator {
	if value != "" {
		if _, err := uuid.Parse(value); err != nil {
			v.errors = append(v.errors, ValidationError{
				Field:   field,
				Message: field + " must be a valid UUID",
			})
		}
	}
	return v
}

// OneOf checks if value is in allowed list
func (v *Validator) OneOf(field, value string, allowed []string) *Validator {
	if value == "" {
		return v
	}
	for _, a := range allowed {
		if value == a {
			return v
		}
	}
	v.errors = append(v.errors, ValidationError{
		Field:   field,
		Message: field + " must be one of: " + strings.Join(allowed, ", "),
	})
	return v
}

// Custom allows custom validation logic
func (v *Validator) Custom(field, message string, condition bool) *Validator {
	if !condition {
		v.errors = append(v.errors, ValidationError{
			Field:   field,
			Message: message,
		})
	}
	return v
}

// IsValid returns true if there are no validation errors
func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

// Errors returns all validation errors
func (v *Validator) Errors() ValidationErrors {
	return v.errors
}

// BindJSON binds JSON from request
func BindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.New("invalid JSON format")
	}
	return nil
}
