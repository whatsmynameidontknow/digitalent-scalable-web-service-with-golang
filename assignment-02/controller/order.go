package controller

import (
	"assignment-02/dto"
	"assignment-02/helper"
	"assignment-02/service"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *orderController {
	return &orderController{orderService}
}

// CreateOrder godoc
// @Summary create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param request body dto.OrderRequest true "required body"
// @Success 201 {object} helper.Response[dto.OrderCreateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /orders [post]
func (c *orderController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.OrderRequest
		resp helper.Response[*dto.OrderCreateResponse] // use pointer so when there's an error, the "data" field will become null instead of a zero-valued struct
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = data.Validate()
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	orderID, err := c.orderService.Create(r.Context(), data)
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusInternalServerError).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = resp.Success(true).Data(&orderID).Code(http.StatusCreated).Send(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllOrders godoc
// @Summary list all orders
// @Tags orders
// @Produce json
// @Success 200 {object} helper.Response[[]dto.OrderResponse]
// @Failure 500 {object} helper.Response[any]
// @Router /orders [get]
func (c *orderController) GetAll(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[[]dto.OrderResponse]
	data, err := c.orderService.GetAll(r.Context())
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusInternalServerError).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = resp.Data(data).Success(true).Code(http.StatusOK).Send(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetOrderByID godoc
// @Summary get a specific order by their ID
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} helper.Response[dto.OrderResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /orders/{orderId} [get]
func (c *orderController) GetByID(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[*dto.OrderResponse]
	orderIDStr := r.PathValue("orderId")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		err = resp.Success(false).Error(errors.New("order id must be >= 0")).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	data, err := c.orderService.GetByID(r.Context(), uint(orderID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("order with id = %d can't be found", orderID)
			err = resp.Success(false).Error(err).Code(http.StatusNotFound).Send(w)
		} else {
			err = resp.Success(false).Error(err).Code(http.StatusInternalServerError).Send(w)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = resp.Data(&data).Success(true).Code(http.StatusOK).Send(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteOrderByID godoc
// @Summary delete order by their ID
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} helper.Response[any]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /orders/{orderId} [delete]
func (c *orderController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp helper.Response[any]
	orderIDStr := r.PathValue("orderId")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		err = resp.Success(false).Error(errors.New("order id must be >= 0")).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = c.orderService.Delete(r.Context(), uint(orderID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("order with id = %d can't be found", orderID)
			err = resp.Success(false).Error(err).Code(http.StatusNotFound).Send(w)
		} else {
			err = resp.Success(false).Error(err).Code(http.StatusInternalServerError).Send(w)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = resp.Success(true).Code(http.StatusOK).Send(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateOrderByID godoc
// @Summary update an order by their ID
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Param request body dto.OrderRequest true "required body"
// @Success 200 {object} helper.Response[any]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /orders/{orderId} [put]
func (c *orderController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		resp helper.Response[any]
		data dto.OrderRequest
	)
	orderIDStr := r.PathValue("orderId")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		err = resp.Success(false).Error(errors.New("order id must be >= 0")).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = data.Validate()
	if err != nil {
		err = resp.Success(false).Error(err).Code(http.StatusBadRequest).Send(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = c.orderService.Update(r.Context(), uint(orderID), data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("order with id = %d can't be found", orderID)
			err = resp.Success(false).Error(err).Code(http.StatusNotFound).Send(w)
		} else {
			err = resp.Success(false).Error(err).Code(http.StatusInternalServerError).Send(w)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = resp.Success(true).Code(http.StatusOK).Send(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
