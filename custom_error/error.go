package CustomError

import "fmt"

// CustomError represents a structured error with code, message, and details
type CustomError struct {
	Code    int
	Message string
	Details string
}

// String returns the string representation of the error
func (e *CustomError) String() string {
	if e.Details != "" {
		return fmt.Sprintf(
			"[%d] %s: %s",
			e.Code,
			e.Message,
			e.Details,
		)
	}
	return fmt.Sprintf(
		"[%d] %s",
		e.Code,
		e.Message,
	)
}

// NewCustomError creates a new CustomError with optional details override
func NewCustomError(
	err *CustomError,
	details string,
) *CustomError {
	return &CustomError{
		Code:    err.Code,
		Message: err.Message,
		Details: details,
	}
}
