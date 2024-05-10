package telegraph_errors

import "fmt"

// ValidationError represents an error due to invalid input parameters.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Message)
}

// APIError represents an error returned by the Telegraph API.
type APIError struct {
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("telegraph API error: %s", e.Message)
}
