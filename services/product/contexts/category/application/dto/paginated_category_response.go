package dto

import "github.com/Akiles94/go-test-api/shared/application/shared_dto"

type PaginatedCategoryResponse = shared_dto.PaginatedResult[CategoryResponse]
