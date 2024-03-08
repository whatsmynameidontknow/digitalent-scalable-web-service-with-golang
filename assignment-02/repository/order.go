package repository

import (
	"assignment-02/model"
	"context"
	"database/sql"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Insert(ctx context.Context, tx *sql.Tx, data model.Order) (uint, error) {
	query := "INSERT INTO orders(customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id"
	row := tx.QueryRowContext(ctx, query, data.CustomerName, data.OrderedAt)
	if err := row.Err(); err != nil {
		return data.ID, err
	}

	if err := row.Scan(&data.ID); err != nil {
		return data.ID, err
	}

	return data.ID, nil
}

func (r *orderRepository) GetAll(ctx context.Context) ([]model.Order, error) {
	query := `SELECT o.order_id, o.customer_name, o.ordered_at, i.item_id, i.item_code, i.description, i.quantity
	FROM orders o
	LEFT JOIN items i ON o.order_id=i.order_id
	ORDER BY o.order_id ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[uint][]model.Item)
	var data []model.Order
	for rows.Next() {
		var order model.Order
		var item model.Item
		rows.Scan(&order.ID, &order.CustomerName, &order.OrderedAt, &item.ID, &item.ItemCode, &item.Description, &item.Quantity)
		if _, ok := m[order.ID]; !ok {
			data = append(data, order)
		}
		m[order.ID] = append(m[order.ID], item)
	}

	for i := range data {
		data[i].Items = m[data[i].ID]
	}

	return data, nil
}

func (r *orderRepository) GetByID(ctx context.Context, id uint) (model.Order, error) {
	var data model.Order

	query := `SELECT o.order_id, o.customer_name, o.ordered_at, i.item_id, i.item_code, i.description, i.quantity
	FROM orders o
	LEFT JOIN items i ON o.order_id=i.order_id
	WHERE o.order_id = $1
	ORDER BY i.quantity ASC`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		rows.Scan(&data.ID, &data.CustomerName, &data.OrderedAt, &item.ID, &item.ItemCode, &item.Description, &item.Quantity)
		data.Items = append(data.Items, item)
	}

	if data.IsEmpty() {
		return data, sql.ErrNoRows
	}

	return data, nil
}

func (r *orderRepository) Delete(ctx context.Context, id uint) error {
	query := "DELETE FROM orders WHERE order_id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *orderRepository) Update(ctx context.Context, tx *sql.Tx, data model.Order) error {
	query := "UPDATE orders SET(customer_name, ordered_at) = ($1, $2) WHERE order_id = $3"
	res, err := tx.ExecContext(ctx, query, data.CustomerName, data.OrderedAt, data.ID)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n == 0 {
		return sql.ErrNoRows
	}

	return nil
}
