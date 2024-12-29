package dto

type ClassResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Leader    int    `json:"leader"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
