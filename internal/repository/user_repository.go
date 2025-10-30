package repository

import (
	"github.com/Dawit0/cleanArch/internal/domain"
	"gorm.io/gorm"
)

type UserRepositery interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepositery.
func (u *userRepository) Create(product *domain.Product) error {
	return u.db.Create(product).Error
}

// DeleteUser implements UserRepositery.
func (u *userRepository) DeleteProduct(id uint) error {
	return u.db.Delete(&domain.Product{}, id).Error
}

// GetAll implements UserRepositery.
func (u *userRepository) GetAll() ([]domain.Product, error) {
	var product []domain.Product
	err := u.db.Find(&product).Error
	return product, err
}

// GetById implements UserRepositery.
func (u *userRepository) GetById(id int) (domain.Product, error) {
	var product domain.Product
	err := u.db.First(&product, id).Error
	return product, err
}

// UpdateUser implements UserRepositery.
func (u *userRepository) UpdateProduct(product *domain.Product) error {
	return u.db.Save(product).Error
}

func NewUserRepository(db *gorm.DB) UserRepositery {
	return &userRepository{db}
}
