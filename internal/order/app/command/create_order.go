package command

import (
	"context"
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
	// TODO: call stock grpc to get items
	err := c.stockGRPC.CheckIfItemsInStock(ctx, cmd.Items)
	resp, err := c.stockGRPC.GetItems(ctx, []string{"123"})
	logrus.Info("createOrderHandler || resp from stockGRPC.GetItems", resp)
	var stockResponse []*orderpb.Item
	for _, item := range cmd.Items {
		stockResponse = append(stockResponse, &orderpb.Item{
			Id:       item.ItemId,
			Quantity: item.Quantity,
		})
	}
	order, err := c.orderRepo.Create(ctx, &domain.Order{
		CustomerId: cmd.CustomerId,
		Items:      stockResponse,
	})

	if err != nil {
		return nil, err
	}

	return &CreateOrderResult{
		OrderId: order.Id,
	}, nil
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
