package dto

import "github.com/Durgarao310/zneha-backend/internal/model"

// ProductCreateRequest represents payload for creating a product
type ProductCreateRequest struct {
	Name             string `json:"name" binding:"required,min=3,max=255"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty" binding:"max=255"`
	Status           string `json:"status,omitempty" binding:"omitempty,oneof=active inactive"`
}

// ProductUpdateRequest represents payload for updating a product (same required fields for simplicity)
// If partial updates are needed later, switch fields to pointer types and adjust logic.
type ProductUpdateRequest struct {
	Name             string `json:"name" binding:"required,min=3,max=255"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty" binding:"max=255"`
	Status           string `json:"status,omitempty" binding:"omitempty,oneof=active inactive"`
}

// ProductResponse represents product data returned to clients
type ProductResponse struct {
	ID               uint64 `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	Status           string `json:"status"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

// ToProductResponse converts model to response DTO
func ToProductResponse(m *model.Product) ProductResponse {
	if m == nil {
		return ProductResponse{}
	}
	return ProductResponse{
		ID:               m.ID,
		Name:             m.Name,
		Description:      m.Description,
		ShortDescription: m.ShortDescription,
		Status:           m.Status,
		CreatedAt:        m.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:        m.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// ToProductResponseList converts slice of models to slice of responses
func ToProductResponseList(list []model.Product) []ProductResponse {
	out := make([]ProductResponse, 0, len(list))
	for i := range list {
		out = append(out, ToProductResponse(&list[i]))
	}
	return out
}
