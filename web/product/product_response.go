package product

import (
	"backend-inventory-app/entity"
	"time"
)

type ProductResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Category  entity.Category `json:"category"`
	HargaBeli int             `json:"harga_beli"`
	Qty       int             `json:"qty"`
	Image     string          `json:"image"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
