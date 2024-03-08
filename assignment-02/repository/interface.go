package repository

import (
	"assignment-02/model"
	"context"
	"database/sql"
)

type OrderRepository interface {
	Insert(context.Context, *sql.Tx, model.Order) (uint, error)
	GetAll(context.Context) ([]model.Order, error)
	GetByID(context.Context, uint) (model.Order, error)
	Delete(context.Context, uint) error
	Update(context.Context, *sql.Tx, model.Order) error
}

type ItemRepository interface {
	InsertMultiple(context.Context, *sql.Tx, []model.Item) error
	DeleteByOrderID(context.Context, *sql.Tx, uint) error
}
