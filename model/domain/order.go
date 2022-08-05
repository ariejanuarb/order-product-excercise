package domain

import "time"

type Orders struct {
	Id            int
	InvoiceNumber string
	OrderDate     time.Time
	CustomerId    int
	Amount        float64
	TotalAmount   float64
	Discount      float64
	TotalDiscount float64
}
