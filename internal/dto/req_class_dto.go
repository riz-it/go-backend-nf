package dto

type CreateClassRequest struct {
	Name     string `json:"name" validate:"required"`
	Leader   int    `json:"leader"`
	IsActive bool   `json:"is_active"`
}
