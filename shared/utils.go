package shared

import "fmt"

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