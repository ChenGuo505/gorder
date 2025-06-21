package command

import (
	"context"
	"errors"
	"github.com/ChenGuo505/gorder/order/app/query"

	"github.com/ChenGuo505/gorder/common/decorator"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	domain "github.com/ChenGuo505/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type CreateOrder struct {
	CustomerId string
	Items      []*orderpb.ItemWithQuantity
}

type CreateOrderResult struct {
	OrderId string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder, *CreateOrderResult]

type createOrderHandler struct {
	orderRepo domain.Repository
	stockGRPC query.StockService
}

func (c createOrderHandler) Handle(ctx context.Context, cmd CreateOrder) (*CreateOrderResult, error) {
	validItems, err := c.validate(ctx, cmd.Items)
	if err != nil {
		return nil, err
	}
	order, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerId: cmd.CustomerId,
		Items:      validItems,
	})

	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{
		OrderId: order.Id,
	}, nil
}

func (c createOrderHandler) validate(ctx context.Context, items []*orderpb.ItemWithQuantity) ([]*orderpb.Item, error) {
	if len(items) == 0 {
		return nil, errors.New("must provide at least one item")
	}
	items = packItems(items)
	resp, err := c.stockGRPC.CheckIfItemsInStock(ctx, items)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func packItems(items []*orderpb.ItemWithQuantity) []*orderpb.ItemWithQuantity {
	merged := make(map[string]int32)
	for _, item := range items {
		merged[item.ItemId] += item.Quantity
	}
	var packedItems []*orderpb.ItemWithQuantity
	for itemId, quantity := range merged {
		packedItems = append(packedItems, &orderpb.ItemWithQuantity{
			ItemId:   itemId,
			Quantity: quantity,
		})
	}
	return packedItems
}

func NewCreateOrderHandler(orderRepo domain.Repository, stockGRPC query.StockService, logger *logrus.Entry, metricsClient decorator.MetricsClient) CreateOrderHandler {
	if orderRepo == nil {
		panic("orderRepo cannot be nil")
	}
	return decorator.AppluCommandDecorators(
		createOrderHandler{
			orderRepo: orderRepo,
			stockGRPC: stockGRPC,
		},
		logger,
		metricsClient,
	)
}
