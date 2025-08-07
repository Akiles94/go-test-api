package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
)

type GetAllProductsUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewGetAllProductsUseCase(repo outbound.ProductRepositoryPort) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		repo: repo,
	}
}

func (uc *GetAllProductsUseCase) Execute(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error) {
	products, nextCursor, err := uc.repo.GetAll(ctx, cursor, limit)
	if err != nil {
		return nil, nil, err
	}

	return products, nextCursor, nil
}
