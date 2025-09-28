package types

type PaginatedResult[T any] struct {
	Data      []T   `json:"data"`
	Total     int64 `json:"total"`
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalPage int   `json:"total_page"`
}

type PaginationSearchParam struct {
	Page      int
	PageSize  int
	Search    string
	SearchKey []string
}

type PaginationParam struct {
	Page     int
	PageSize int
}
