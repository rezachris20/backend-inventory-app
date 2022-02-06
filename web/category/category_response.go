package category

import (
	"time"
)

type CategoryResponse struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	IsPrimary   bool        `json:"is_primary"`
	IsActive    bool        `json:"is_active"`
	CategoryID  int         `json:"category_id"`
	SubCategory interface{} `json:"sub_category"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type SubCategoryResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	IsPrimary  bool      `json:"is_primary"`
	IsActive   bool      `json:"is_active"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
