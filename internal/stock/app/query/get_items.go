package query

import (
	"context"
	"github.com/ChenGuo505/gorder/common/decorator"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	domain "github.com/ChenGuo505/gorder/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type GetItems struct {
	ItemIds []string
}

type GetItemsHandler decorator.QueryHandler[GetItems, []*orderpb.Item]

type getItemsHandler struct {
	stockRepo domain.Repository
}

func (g getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*orderpb.Item, error) {
	items, err := g.stockRepo.GetItems(ctx, query.ItemIds)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func NewGetItemHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) GetItemsHandler {
	if stockRepo == nil {
		panic("stockRepo cannot be nil")
	}
	return decorator.AppluQueryDecorators(
		getItemsHandler{stockRepo: stockRepo},
		logger,
		metricClient,
	)
}
