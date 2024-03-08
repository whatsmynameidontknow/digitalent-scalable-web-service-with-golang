package model

import "time"

type Order struct {
	ID           uint
	CustomerName string
	OrderedAt    time.Time
	Items        []Item
}

func (o Order) IsEmpty() bool {
	return o.ID == 0 && o.CustomerName == "" && o.OrderedAt.IsZero() && o.Items == nil
}
