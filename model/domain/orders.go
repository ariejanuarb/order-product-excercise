package domain

type Orders struct {
	Id            int
	InvoiceNo     string
	OrderDate     string
	CustomerId    int
	Amount        int
	TotalAmount   int
	Discount      int
	TotalDiscount int
}
