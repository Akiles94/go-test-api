package dto

type PaginatedResult[T any] struct {
	Items      []T     `json:"items"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

func NewPaginatedResult[T any](items []T, nextCursor *string) PaginatedResult[T] {
	return PaginatedResult[T]{
		Items:      items,
		NextCursor: nextCursor,
	}
}

func (pr *PaginatedResult[T]) HasMore() bool {
	return pr.NextCursor != nil
}

func (pr *PaginatedResult[T]) Count() int {
	return len(pr.Items)
}
