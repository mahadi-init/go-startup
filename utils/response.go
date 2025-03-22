package utils

// Response is a standardized API response
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse returns a standardized success response
func SuccessResponse(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// ErrorResponse returns a standardized error response
func ErrorResponse(err string, message string) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}
