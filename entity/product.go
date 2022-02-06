package entity

import "time"

type Product struct {
	ID         int
	Name       string
	CategoryID int
	Category   Category
	HargaBeli  int
	Qty        int
	Image      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
