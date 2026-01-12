package utils

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// AppError represents an application-level error
type AppError struct {
	Message    string
	StatusCode int
	Err        error
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError
func NewAppError(message string, statusCode int, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

// IsNotFoundError checks if the error is a "not found" error
func IsNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// IsDuplicateError checks if the error is a duplicate key error
func IsDuplicateError(err error) bool {
	if err == nil {
		return false
	}
	// PostgreSQL duplicate key error code is 23505
	return errors.Is(err, gorm.ErrDuplicatedKey)
}

// WrapError wraps an error with additional context
func WrapError(message string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// HandleDatabaseError converts database errors to appropriate AppError
func HandleDatabaseError(err error, resourceName string) *AppError {
	if IsNotFoundError(err) {
		return NewAppError(
			fmt.Sprintf("%s not found", resourceName),
			404,
			err,
		)
	}
	
	if IsDuplicateError(err) {
		return NewAppError(
			fmt.Sprintf("%s already exists", resourceName),
			409,
			err,
		)
	}
	
	return NewAppError(
		fmt.Sprintf("Database error while processing %s", resourceName),
		500,
		err,
	)
}
