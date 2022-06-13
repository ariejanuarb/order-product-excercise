package domain

type OrdersProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     int
	Amount    int
}
