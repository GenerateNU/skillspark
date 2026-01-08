// Package xvalidator provides validation utilities for struct validation using go-playground/validator
package xvalidator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ErrorResponse represents a validation error with detailed information
type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Param       string
	Value       interface{}
}

// XValidator wraps the validator instance for struct validation
type XValidator struct {
	Validator *validator.Validate
}

// GlobalErrorHandlerResp represents a global error response structure
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Validate is the global validator instance
var Validate = validator.New()

// Validate validates the given struct and returns a slice of validation errors
func (v XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := Validate.Struct(data)
	if errs != nil {
		// Use errors.As to safely check for ValidationErrors
		var validationErrs validator.ValidationErrors
		if errors.As(errs, &validationErrs) {
			for _, err := range validationErrs {
				// In this case data holds the struct being validated
				var elem ErrorResponse
				elem.FailedField = err.Field() // Export struct field name
				elem.Tag = err.Tag()           // Export struct tag
				elem.Param = err.Param()       // Export tag parameter
				elem.Value = err.Value()       // Export field value
				elem.Error = true

				if param := err.Param(); param != "" {
					elem.Tag = fmt.Sprintf("%s:%s", elem.Tag, param)
				}

				validationErrors = append(validationErrors, elem)
			}
		}
	}

	return validationErrors
}

// ConvertToMessages converts validation errors to user-friendly messages
func ConvertToMessages(errors []ErrorResponse) map[string]string {
	errorMap := make(map[string]string)

	for _, err := range errors {
		field := strings.ToLower(err.FailedField)
		message := getErrorMessage(err)
		errorMap[field] = message
		// slog.Error(message)
	}

	return errorMap
}

func extractOtherField(tag string) string {
	parts := strings.Split(tag, ":")
	if len(parts) == 2 {
		return parts[1]
	}
	return "the referenced field"
}

// getErrorMessage generates user-friendly error messages
func getErrorMessage(err ErrorResponse) string {
	field := err.FailedField
	tag := err.Tag

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s does not meet the minimum allowed value or length", field)
	case "max":
		return fmt.Sprintf("%s exceeds the maximum allowed value or length", field)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", field, err.Param)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %v", field, err.Param)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "gtfield":
		return fmt.Sprintf("%s must be greater than %s", field, extractOtherField(tag))
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// Validator is the global XValidator instance
var Validator = &XValidator{
	Validator: Validate,
}
