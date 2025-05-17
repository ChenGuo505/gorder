package query

import (
	"context"

	"github.com/ChenGuo505/gorder/common/decorator"
	domain "github.com/ChenGuo505/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type GetCustomerOrder struct {
	CustomerId string
	OrderId    string
}

type GetCustomerOrderHandler decorator.QueryHandler[GetCustomerOrder, *domain.Order]

type getCustomerOrderHandler struct {
	orderRepo domain.Repository
}

func (g getCustomerOrderHandler) Handle(ctx context.Context, cmd GetCustomerOrder) (*domain.Order, error) {
	order, err := g.orderRepo.Get(ctx, cmd.OrderId, cmd.CustomerId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func NewGetCustomerOrderHandler(orderRepo domain.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) GetCustomerOrderHandler {
	if orderRepo == nil {
		panic("orderRepo cannot be nil")
	}
	return decorator.AppluQueryDecorators(
		getCustomerOrderHandler{
			orderRepo: orderRepo,
		},
		logger,
		metricsClient,
	)
}
