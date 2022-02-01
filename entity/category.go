package entity

import "time"

type Category struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"size:255;index:idx_name,unique"`
	IsPrimary  bool
	IsActive   bool
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
