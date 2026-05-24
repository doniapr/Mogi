package users

import (
	"mogi/internal/domain"
	"mogi/internal/infrastructure/postgres"
)

// UseCase handles business logic for User
type UseCase struct {
	repo postgres.Repository
}

// NewUseCase creates new usecase instance
func NewUseCase(repo postgres.Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

// GetAll retrieves all users
func (uc *UseCase) GetAll() ([]domain.User, error) {
	return uc.repo.FindAll()
}

// GetByID retrieves a user by ID
func (uc *UseCase) GetByID(id int) (*domain.User, error) {
	return uc.repo.FindByID(id)
}

// Create creates a new user
func (uc *UseCase) Create(user *domain.User) error {
	return uc.repo.Create(user)
}

// Update updates a user
func (uc *UseCase) Update(user *domain.User) error {
	return uc.repo.Update(user)
}

// Delete deletes a user
func (uc *UseCase) Delete(id int) error {
	return uc.repo.Delete(id)
}
