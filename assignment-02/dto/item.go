package dto

import (
	"errors"
	"fmt"
)

type ItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (i ItemRequest) Validate() error {
	var err error

	if i.ItemCode == "" {
		err = errors.Join(err, errors.New("item_code can't be empty"))
		return err // return immediately, biar kalo quantity nya < 1, item_code = %s nya ga kosong. wkwk
	}

	if i.Quantity < 1 {
		err = errors.Join(err, fmt.Errorf("item_code = %s. quantity can't be less than 1", i.ItemCode))
	}

	return err
}

type ItemResponse struct {
	ID          uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
