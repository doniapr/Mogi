package dto

import "time"

// CreateUserRequest is the DTO for creating a new User
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserRequest is the DTO for updating a User
type UpdateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// CreateUserResponse is the DTO for create User response
type CreateUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// UserResponse is the DTO for User response
type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllUserResponse is the DTO for get all User response
type GetAllUserResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    []UserResponse  `json:"data"`
}

// ErrorResponse is the DTO for error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

