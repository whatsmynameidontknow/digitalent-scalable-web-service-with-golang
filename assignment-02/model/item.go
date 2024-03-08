package model

type Item struct {
	ID, OrderID           uint
	Quantity              int
	ItemCode, Description string
}
