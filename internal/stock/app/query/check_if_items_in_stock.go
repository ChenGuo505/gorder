package query

import (
	"context"
	"github.com/ChenGuo505/gorder/common/decorator"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	domain "github.com/ChenGuo505/gorder/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type CheckIfItemsInStock struct {
	Items []*orderpb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*orderpb.Item]

type checkIfItemsInStockHandler struct {
	stockRepo domain.Repository
}

func (c checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*orderpb.Item, error) {
	var (
		ids          []string
		idToQuantity = make(map[string]int32)
	)
	for _, q := range query.Items {
		ids = append(ids, q.ItemId)
		idToQuantity[q.ItemId] = q.Quantity
	}

	itemsInStock, err := c.stockRepo.GetItems(ctx, ids)
	if err != nil {
		return nil, err
	}

	var res []*orderpb.Item
	for _, item := range itemsInStock {
		need, ok := idToQuantity[item.Id]
		if !ok {
			continue
		}
		if item.Quantity >= need {
			res = append(res, &orderpb.Item{
				Id:       item.Id,
				Name:     item.Name,
				Quantity: item.Quantity,
				PriceId:  item.PriceId,
			})
		}
	}
	return res, nil
}

func NewCheckIfItemsInStockHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	if stockRepo == nil {
		panic("stockRepo cannot be nil")
	}
	return decorator.AppluQueryDecorators(
		checkIfItemsInStockHandler{stockRepo: stockRepo},
		logger,
		metricClient,
	)
}
