package handler

import (
	"mogi/internal/domain"
	"mogi/internal/dto"
	"mogi/internal/usecase/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler handles HTTP requests for User
type UserHandler struct {
	uc *users.UseCase
}

// NewUserHandler creates new handler instance
func NewUserHandler(uc *users.UseCase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

// GetAll handles GET /users endpoint
func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.uc.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to retrieve users",
			Error:   err.Error(),
		})
	}

	// Convert to response DTOs
	responses := make([]dto.UserResponse, len(users))
	for i, u := range users {
		responses[i] = dto.UserResponse{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, dto.GetAllUserResponse{
		Success: true,
		Message: "Success",
		Data:    responses,
	})
}

// GetByID handles GET /users/:id endpoint
func (h *UserHandler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
	}

	user, err := h.uc.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Success: false,
			Message: "User not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// Create handles POST /users endpoint
func (h *UserHandler) Create(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	// Basic validation
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Email is required",
		})
	}

	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Password is required",
		})
	}

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.uc.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.CreateUserResponse{
		Success: true,
		Message: "User created successfully",
	})
}

// Update handles PUT /users/:id endpoint
func (h *UserHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	user := &domain.User{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
	}

	err = h.uc.Update(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CreateUserResponse{
		Success: true,
		Message: "User updated successfully",
	})
}

// Delete handles DELETE /users/:id endpoint
func (h *UserHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
	}

	err = h.uc.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CreateUserResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
