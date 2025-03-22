package models

// User represents a user model with necessary fields.
type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"password"`
	Age      *int    `json:"age,omitempty"`
	Gender   *string `json:"gender,omitempty"`
}
