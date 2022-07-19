package web

type OrdersCreateRequest struct {
	InvoiceNumber string                 `validate:"required,min=1,max=100" json:"invoice_number"`
	OrderDate     int                    `validate:"required" json:"order_date"`
	CustomerId    int                    `validate:"required" json:"customer_id"`
	Amount        int                    `validate:"required" json:"amount"`
	TotalAmount   int                    `validate:"required" json:"total_amount"`
	Discount      int                    `validate:"required" json:"discount"`
	TotalDiscount int                    `validate:"required" json:"total_discount"`
	Details       []OrdersProductRequest `validate:"required" json:"details"`
}

type OrdersProductRequest struct {
	ProductId int `validate:"required" json:"product_id"`
	Qty       int `validate:"required" json:"qty"`
	Price     int `validate:"required" json:"price"`
	Amount    int `validate:"required" json:"amount"`
}
