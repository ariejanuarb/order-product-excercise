package web

import "time"

type OrderCreateRequest struct {
	InvoiceNumber string                `validate:"required,min=1,max=100" json:"invoice_number"`
	OrderDate     time.Time             `validate:"required" json:"order_date"`
	CustomerId    int                   `validate:"required" json:"customer_id"`
	Amount        float64               `validate:"required" json:"amount"`
	TotalAmount   float64               `validate:"required" json:"total_amount"`
	Discount      float64               `validate:"required" json:"discount"`
	TotalDiscount float64               `validate:"required" json:"total_discount"`
	Details       []OrderProductRequest `validate:"required" json:"details"`
}

type OrderProductRequest struct {
	OrderId   int `validate:"required" json:"order_id"`
	ProductId int `validate:"required" json:"product_id"`
	Qty       int `validate:"required" json:"qty"`
	Price     int `validate:"required" json:"price"`
	Amount    int `validate:"required" json:"amount"`
}
