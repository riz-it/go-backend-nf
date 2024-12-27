package dto

type ApiResponse[T any] struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    *T            `json:"data"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  any           `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
