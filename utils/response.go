package utils

// Response is a standardized API response
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse returns a standardized success response
func SuccessResponse(data any, message string) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse returns a standardized error response
func ErrorResponse(err string) Response {
	return Response{
		Success: false,
		Error:   err,
	}
}
