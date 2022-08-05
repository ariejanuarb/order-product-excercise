package web

import "time"

type OrderUpdateRequest struct {
	Id            int                  `validate:"required" json:"id"`
	InvoiceNumber string               `validate:"required,min=1,max=100" json:"invoice_number"`
	OrderDate     time.Time            `validate:"required" json:"order_date"`
	CustomerId    int                  `validate:"required" json:"customer_id"`
	Amount        int                  `validate:"required" json:"amount"`
	TotalAmount   int                  `validate:"required" json:"total_amount"`
	Discount      int                  `validate:"required" json:"discount"`
	TotalDiscount int                  `validate:"required" json:"total_discount"`
	Details       []OrderCreateRequest `validate:"required" json:"details"`
}
