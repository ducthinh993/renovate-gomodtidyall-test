package shared

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// Helper function that will be used across modules
func FormatMessage(message string) string {
	return fmt.Sprintf("[SHARED] %s", message)
}

// Common data structure
type Response struct {
	Status  string
	Message string
}

func NewResponse(status, message string) Response {
	return Response{
		Status:  status,
		Message: FormatMessage(message),
	}
}

// ValidateResponse uses testify to validate response format
func ValidateResponse(resp Response) bool {
	assert.NotEmpty(nil, resp.Status, "Status should not be empty")
	assert.NotEmpty(nil, resp.Message, "Message should not be empty")
	return resp.Status != "" && resp.Message != ""
}