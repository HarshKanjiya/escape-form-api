package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is a global validator instance
var Validator = validator.New()

// ValidateStruct validates a struct and returns validation errors
func ValidateStruct(s interface{}) error {
	return Validator.Struct(s)
}

// ValidationErrorResponse represents a validation error response
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ParseValidationErrors parses validator errors into a readable format
func ParseValidationErrors(err error) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errors = append(errors, ValidationErrorResponse{
				Field:   getJSONFieldName(fieldError),
				Message: getErrorMessage(fieldError),
			})
		}
	}

	return errors
}

// getJSONFieldName extracts the JSON field name from validation error
func getJSONFieldName(fieldError validator.FieldError) string {
	field := fieldError.Field()
	// Convert to snake_case for JSON compatibility
	return toSnakeCase(field)
}

// getErrorMessage generates a human-readable error message
func getErrorMessage(fieldError validator.FieldError) string {
	field := fieldError.Field()

	switch fieldError.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", field, fieldError.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, fieldError.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, fieldError.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, fieldError.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, fieldError.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, fieldError.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// toSnakeCase converts camelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// ValidateRequest validates request body and returns error response if validation fails
// func ValidateRequest(c *fiber.Ctx, payload interface{}) error {
// 	// Parse body into the payload struct
// 	if err := c.BodyParser(payload); err != nil {
// 		return BadRequest(c, "Invalid request body")
// 	}

// 	// Validate the struct
// 	if err := ValidateStruct(payload); err != nil {
// 		validationErrors := ParseValidationErrors(err)
// 		return ValidationError(c, validationErrors)
// 	}

// 	return nil
// }

// GetStructTag returns the value of a struct tag
func GetStructTag(s interface{}, fieldName, tagName string) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, found := t.FieldByName(fieldName)
	if !found {
		return ""
	}

	return field.Tag.Get(tagName)
}
