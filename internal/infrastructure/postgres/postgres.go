package postgres

import (
	"mogi/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database represents the database connection
type Database struct {
	DB *gorm.DB
}

// NewDatabase initializes a new database connection
func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migration
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

// Repository interface for User
type Repository interface {
	FindAll() ([]domain.User, error)
	FindByID(id int) (*domain.User, error)
	Create(data *domain.User) error
	Update(data *domain.User) error
	Delete(id int) error
}

// UserRepository implements Repository interface
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates new repository instance
func NewUserRepository(database *Database) Repository {
	return &UserRepository{
		db: database.DB,
	}
}

// FindAll retrieves all User records
func (r *UserRepository) FindAll() ([]domain.User, error) {
	var Users []domain.User
	err := r.db.Find(&Users).Error
	return Users, err
}

// FindByID retrieves a User by ID
func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	var User domain.User
	err := r.db.Where("id = ?", id).First(&User).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

// Create creates a new User record
func (r *UserRepository) Create(data *domain.User) error {
	return r.db.Create(data).Error
}

// Update updates a User record
func (r *UserRepository) Update(data *domain.User) error {
	return r.db.Save(data).Error
}

// Delete deletes a User record
func (r *UserRepository) Delete(id int) error {
	return r.db.Delete(&domain.User{}, id).Error
}
