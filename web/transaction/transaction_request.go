package transaction

type CreateTransactionRequest struct {
	Diskon             int                  `json:"diskon"`
	TotalBayar         int                  `json:"total_bayar" validate:"required"`
	DetailTransactions []DetailTransactions `json:"detail_transactions" validate:"required"`
}

type DetailTransactions struct {
	ProductID   int `json:"product_id" validate:"required"`
	Qty         int `json:"qty" validate:"required"`
	HargaSatuan int `json:"harga_satuan" validate:"required"`
	Total       int `json:"total" validate:"required"`
}
