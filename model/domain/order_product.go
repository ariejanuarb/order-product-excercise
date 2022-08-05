package domain

type OrdersProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     float64
	Amount    float64
}
