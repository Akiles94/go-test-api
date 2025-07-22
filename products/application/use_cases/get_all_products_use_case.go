package use_cases

import (
	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
	"github.com/Akiles94/go-test-api/products/domain/models"
)

type GetAllProductsUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewGetAllProductsUseCase(repo outbound.ProductRepositoryPort) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		repo: repo,
	}
}

func (uc *GetAllProductsUseCase) Execute(cursor *string, limit *int) ([]models.Product, *string, error) {
	products, nextCursor, err := uc.repo.GetAll(cursor, limit)
	if err != nil {
		return nil, nil, err
	}

	return products, nextCursor, nil
}
