package pagination

const Key = "pagination"

type Meta struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

type Result[T any] struct {
	Data []T  `json:"data"`
	Meta Meta `json:"meta"`
}
