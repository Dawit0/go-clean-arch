package usecase

import (
	"github.com/Dawit0/cleanArch/internal/domain"
	"github.com/Dawit0/cleanArch/internal/repository"
)

type ProductUsecase interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
}

type productUsecase struct {
	repo repository.UserRepositery
}

// Create implements ProductUsecase.
func (p *productUsecase) Create(product *domain.Product) error {
	return p.repo.Create(product)
}

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(id uint) error {
	return p.repo.DeleteProduct(id)
}

// GetAll implements ProductUsecase.
func (p *productUsecase) GetAll() ([]domain.Product, error) {
	return p.repo.GetAll()
}

// GetById implements ProductUsecase.
func (p *productUsecase) GetById(id int) (domain.Product, error) {
	return p.repo.GetById(id)
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(product *domain.Product) error {
	return p.repo.UpdateProduct(product)
}

func NewUsecase(repos repository.UserRepositery) ProductUsecase {
	return &productUsecase{repos}
}
