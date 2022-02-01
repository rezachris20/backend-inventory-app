package category

type CreateCategoryRequest struct {
	Name       string `json:"name" validate:"required"`
	IsPrimary  bool   `json:"is_primary"`
	IsActive   bool   `json:"is_active"`
	CategoryID int    `json:"category_id"`
}

type UpdateCategoryRequest struct {
	Name       string `json:"name"`
	IsPrimary  bool   `json:"is_primary"`
	IsActive   bool   `json:"is_active"`
	CategoryID int    `json:"category_id"`
}
