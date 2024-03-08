package service

import (
	"assignment-02/dto"
	"context"
)

type OrderService interface {
	Create(context.Context, dto.OrderRequest) (dto.OrderCreateResponse, error)
	GetAll(context.Context) ([]dto.OrderResponse, error)
	GetByID(context.Context, uint) (dto.OrderResponse, error)
	Delete(context.Context, uint) error
	Update(context.Context, uint, dto.OrderRequest) error
}
