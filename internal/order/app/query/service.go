package query

import (
	"context"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
)

type StockService interface {
	CheckIfItemsInStock(context.Context, []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error)
	GetItems(context.Context, []string) ([]*orderpb.Item, error)
}
