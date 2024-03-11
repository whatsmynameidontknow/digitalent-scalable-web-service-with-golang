package dto

import (
	"errors"
	"time"
)

type OrderRequest struct {
	OrderedAt    time.Time     `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	CustomerName string        `json:"customerName" example:"John Doe"`
	Items        []ItemRequest `json:"items"`
}

func (r OrderRequest) Validate() error {
	var err error

	if r.OrderedAt.IsZero() {
		err = errors.Join(err, errors.New("orderedAt can't be empty"))
	}
	if r.CustomerName == "" {
		err = errors.Join(err, errors.New("customerName can't be empty"))
	}
	if len(r.Items) == 0 {
		err = errors.Join(err, errors.New("items can't be empty"))
	}

	for i := range r.Items {
		if itemErr := r.Items[i].Validate(); itemErr != nil {
			err = errors.Join(err, itemErr)
		}
	}

	return err
}

type OrderResponse struct {
	ID           uint           `json:"order_id"`
	OrderedAt    time.Time      `json:"ordered_at"`
	CustomerName string         `json:"customer_name"`
	Items        []ItemResponse `json:"items"`
}

type OrderCreateResponse struct {
	ID uint `json:"order_id"`
}
