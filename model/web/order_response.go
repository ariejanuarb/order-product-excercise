package web

import "time"

type OrderResponse struct {
	Id            int                    `json:"id"`
	InvoiceNumber string                 `json:"invoice_number"`
	OrderDate     time.Time              `json:"order_date"`
	CustomerId    int                    `json:"customer_id"`
	Amount        int                    `json:"amount"`
	TotalAmount   int                    `json:"total_amount"`
	Discount      int                    `json:"discount"`
	TotalDiscount int                    `json:"total_discount"`
	Details       []OrderProductResponse `json:"details"`
}

type OrderProductResponse struct {
	Id        int     `json:"id"`
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Qty       int     `json:"qty"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
}
