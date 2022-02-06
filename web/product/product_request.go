package product

type CreateProductRequest struct {
	Name       string `form:"name" validate:"required"`
	CategoryID int    `form:"category_id" validate:"required"`
	HargaBeli  int    `form:"harga_beli" validate:"required"`
	Qty        int    `form:"qty" validate:"required"`
	Image      string `form:"image"`
}

type UpdateProductRequest struct {
	Name       string `form:"name"`
	CategoryID int    `form:"category_id"`
	HargaBeli  int    `form:"harga_beli"`
	Qty        int    `form:"qty"`
	Image      string `form:"image"`
}
