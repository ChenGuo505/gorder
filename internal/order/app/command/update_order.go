package command

import (
	"context"

	"github.com/ChenGuo505/gorder/common/decorator"
	domain "github.com/ChenGuo505/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	Order    *domain.Order
	UpdateFn func(context.Context, *domain.Order) (*domain.Order, error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder, any]

type updateOrderHandler struct {
	orderRepo domain.Repository
}

func (u updateOrderHandler) Handle(ctx context.Context, cmd UpdateOrder) (any, error) {
	if cmd.UpdateFn == nil {
		logrus.Warnf("UpdateFn is nil, using default update function, order: %#v", cmd.Order)
		cmd.UpdateFn = func(ctx context.Context, order *domain.Order) (*domain.Order, error) {
			return order, nil
		}
	}

	err := u.orderRepo.Update(ctx, cmd.Order, cmd.UpdateFn)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewUpdateOrderHandler(orderRepo domain.Repository, logger *logrus.Entry, metricsClient decorator.MetricsClient) UpdateOrderHandler {
	if orderRepo == nil {
		panic("orderRepo cannot be nil")
	}
	return decorator.AppluCommandDecorators(
		updateOrderHandler{
			orderRepo: orderRepo,
		},
		logger,
		metricsClient,
	)
}
