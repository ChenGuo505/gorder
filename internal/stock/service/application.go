package service

import (
	"context"
	"github.com/ChenGuo505/gorder/common/metrics"
	"github.com/ChenGuo505/gorder/stock/adapters"
	"github.com/ChenGuo505/gorder/stock/app/query"
	"github.com/sirupsen/logrus"

	"github.com/ChenGuo505/gorder/stock/app"
)

func NewApplication(ctx context.Context) app.Application {
	stockRepo := adapters.NewMemoryStockRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, logger, metricsClient),
			GetItems:            query.NewGetItemHandler(stockRepo, logger, metricsClient),
		},
	}
}
