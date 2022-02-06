package entity

import (
	"time"
)

type Category struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:255;index:idx_name,unique" json:"name"`
	IsPrimary   bool       `json:"is_primary"`
	IsActive    bool       `json:"is_active"`
	CategoryID  int        `json:"category_id"`
	SubCategory []Category `json:"sub_category" gorm:"foreignkey:CategoryID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
