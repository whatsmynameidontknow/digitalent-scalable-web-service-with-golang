package repository

import (
	"assignment-02/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type itemRepository struct {
}

func NewItemRepository() *itemRepository {
	return &itemRepository{}
}

func (r *itemRepository) InsertMultiple(ctx context.Context, tx *sql.Tx, data []model.Item) error {
	var query strings.Builder
	query.WriteString("INSERT INTO items(item_code, description, quantity, order_id) VALUES ")

	values := make([]any, 0, len(data)*4) // the struct has 4 fillable fields
	for i := range data {
		query.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		values = append(values, data[i].ItemCode, data[i].Description, data[i].Quantity, data[i].OrderID)
		if i < len(data)-1 {
			query.WriteString(", ")
		}
	}

	res, err := tx.ExecContext(ctx, query.String(), values...)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if int(n) != len(data) {
		return errors.New("ada yang kurang")
	}

	return nil
}

func (r *itemRepository) DeleteByOrderID(ctx context.Context, tx *sql.Tx, orderID uint) error {
	query := "DELETE FROM items WHERE order_id = $1"

	_, err := tx.ExecContext(ctx, query, orderID)
	if err != nil {
		return err
	}

	return nil
}
