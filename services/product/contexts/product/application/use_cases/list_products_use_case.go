package use_cases

import (
	"context"

	"github.com/Akiles94/go-test-api/services/product/contexts/product/application/ports/outbound"
	"github.com/Akiles94/go-test-api/services/product/contexts/product/domain/models"
)

type ListProductsUseCase struct {
	repo outbound.ProductRepositoryPort
}

func NewListProductsUseCase(repo outbound.ProductRepositoryPort) *ListProductsUseCase {
	return &ListProductsUseCase{
		repo: repo,
	}
}

func (uc *ListProductsUseCase) Execute(ctx context.Context, cursor *string, limit *int) ([]models.Product, *string, error) {
	products, nextCursor, err := uc.repo.GetAll(ctx, cursor, limit)
	if err != nil {
		return nil, nil, err
	}

	return products, nextCursor, nil
}
