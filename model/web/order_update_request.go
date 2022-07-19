package web

type OrdersUpdateRequest struct {
	Id            int                   `validate:"required" json:"id"`
	InvoiceNumber string                `validate:"required,min=1,max=100" json:"invoice_number"`
	OrderDate     int                   `validate:"required" json:"order_date"`
	CustomerId    int                   `validate:"required" json:"customer_id"`
	Amount        int                   `validate:"required" json:"amount"`
	TotalAmount   int                   `validate:"required" json:"total_amount"`
	Discount      int                   `validate:"required" json:"discount"`
	TotalDiscount int                   `validate:"required" json:"total_discount"`
	Details       []OrdersCreateRequest `validate:"required" json:"details"`
}
