package web

type ProductSellRequest struct {
	productId string `validate:"required" json:"product_id"`
	Price     int    `validate:"required" json:"price"`
	Qty       int    `validate:"required" json:"qty"`
}
