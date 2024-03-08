package service

import (
	"assignment-02/dto"
	"assignment-02/model"
	"assignment-02/repository"
	"context"
	"database/sql"
)

type orderService struct {
	db              *sql.DB
	orderRepository repository.OrderRepository
	itemRepository  repository.ItemRepository
}

func NewOrderService(db *sql.DB, orderRepository repository.OrderRepository, itemRepository repository.ItemRepository) *orderService {
	return &orderService{db, orderRepository, itemRepository}
}

func (s *orderService) Create(ctx context.Context, data dto.OrderRequest) (dto.OrderCreateResponse, error) {
	var (
		order model.Order
		resp  dto.OrderCreateResponse
	)

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return resp, err
	}

	order.OrderedAt = data.OrderedAt
	order.CustomerName = data.CustomerName

	order.ID, err = s.orderRepository.Insert(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	order.Items = make([]model.Item, len(data.Items))

	for i := range data.Items {
		order.Items[i] = model.Item{
			OrderID:     order.ID,
			ItemCode:    data.Items[i].ItemCode,
			Description: data.Items[i].Description,
			Quantity:    data.Items[i].Quantity,
		}
	}

	err = s.itemRepository.InsertMultiple(ctx, tx, order.Items)
	if err != nil {
		tx.Rollback()
		return resp, err
	}

	tx.Commit()
	resp.ID = order.ID
	return resp, nil
}

func (s *orderService) GetAll(ctx context.Context) ([]dto.OrderResponse, error) {
	data, err := s.orderRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]dto.OrderResponse, 0, len(data))
	for i := range data {
		var order dto.OrderResponse
		order.ID = data[i].ID
		order.CustomerName = data[i].CustomerName
		order.OrderedAt = data[i].OrderedAt
		order.Items = make([]dto.ItemResponse, len(data[i].Items))
		for j := range data[i].Items {
			order.Items[j] = dto.ItemResponse{
				ID:          data[i].Items[j].ID,
				ItemCode:    data[i].Items[j].ItemCode,
				Description: data[i].Items[j].Description,
				Quantity:    data[i].Items[j].Quantity,
			}
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *orderService) GetByID(ctx context.Context, id uint) (dto.OrderResponse, error) {
	var order dto.OrderResponse

	data, err := s.orderRepository.GetByID(ctx, id)
	if err != nil {
		return order, err
	}

	order.ID = data.ID
	order.CustomerName = data.CustomerName
	order.OrderedAt = data.OrderedAt
	order.Items = make([]dto.ItemResponse, len(data.Items))
	for i := range data.Items {
		order.Items[i] = dto.ItemResponse{
			ID:          data.Items[i].ID,
			ItemCode:    data.Items[i].ItemCode,
			Description: data.Items[i].Description,
			Quantity:    data.Items[i].Quantity,
		}
	}

	return order, nil
}

func (s *orderService) Delete(ctx context.Context, id uint) error {
	err := s.orderRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) Update(ctx context.Context, id uint, data dto.OrderRequest) error {
	var order model.Order

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	order.ID = id
	order.CustomerName = data.CustomerName
	order.OrderedAt = data.OrderedAt

	err = s.orderRepository.Update(ctx, tx, order) // update the order first, so if order with specified ID doesn't exist, we can return immediately
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.itemRepository.DeleteByOrderID(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	order.Items = make([]model.Item, len(data.Items))
	for i := range data.Items {
		order.Items[i] = model.Item{
			ItemCode:    data.Items[i].ItemCode,
			Description: data.Items[i].Description,
			Quantity:    data.Items[i].Quantity,
			OrderID:     id,
		}
	}

	err = s.itemRepository.InsertMultiple(ctx, tx, order.Items)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
