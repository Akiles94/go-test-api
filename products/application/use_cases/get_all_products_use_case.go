package use_cases

import (
	"github.com/Akiles94/go-test-api/products/application/dto"
	"github.com/Akiles94/go-test-api/products/application/ports/outbound"
)

type GetAllProductsUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewGetAllProductsUseCase(repo outbound.ProductRepositoryPort) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		repo: repo,
	}
}

func (uc *GetAllProductsUseCase) Execute(cursor *string, limit *int) (*dto.ProductsResponse, error) {
	return uc.repo.GetAll(cursor, limit)
}
