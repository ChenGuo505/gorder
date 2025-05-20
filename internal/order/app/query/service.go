package query

import (
	"context"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
)

type StockService interface {
	CheckIfItemsInStock(context.Context, []*orderpb.ItemWithQuantity) error
	GetItems(context.Context, []string) ([]*orderpb.Item, error)
}
