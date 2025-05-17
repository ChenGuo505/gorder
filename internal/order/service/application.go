package service

import (
	"context"

	metrics "github.com/ChenGuo505/gorder/common/metrics"
	"github.com/ChenGuo505/gorder/order/adapters"
	"github.com/ChenGuo505/gorder/order/app"
	"github.com/ChenGuo505/gorder/order/app/command"
	"github.com/ChenGuo505/gorder/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, logger, metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricsClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient),
		},
	}
}
