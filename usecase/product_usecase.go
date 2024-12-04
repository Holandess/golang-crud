package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	ProductRepository repository.ProductRepository
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
