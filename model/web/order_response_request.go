package web

type OrdersResponse struct {
	Id            int                     `json:"id"`
	InvoiceNumber string                  `json:"invoice_number"`
	OrderDate     int                     `json:"order_date"`
	CustomerId    int                     `json:"customer_id"`
	Amount        int                     `json:"amount"`
	TotalAmount   int                     `json:"total_amount"`
	Discount      int                     `json:"discount"`
	TotalDiscount int                     `json:"total_discount"`
	Details       []OrdersProductResponse `json:"details"`
}

type OrdersProductResponse struct {
	Id        int `json:"id"`
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
	Amount    int `json:"amount"`
}
