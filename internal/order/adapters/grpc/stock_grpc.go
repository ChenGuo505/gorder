package grpc

import (
	"context"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
	"github.com/sirupsen/logrus"
)

type StockGRPC struct {
	client stockpb.StockServiceClient
}

func NewStockGRPC(client stockpb.StockServiceClient) *StockGRPC {
	return &StockGRPC{
		client: client,
	}
}

func (s StockGRPC) CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error) {
	resp, err := s.client.CheckIfItemsInStock(ctx, &stockpb.CheckIfItemsInStockRequest{Items: items})
	logrus.Info("stock_grpc resp", resp)
	return resp, err
}

func (s StockGRPC) GetItems(ctx context.Context, itemsId []string) ([]*orderpb.Item, error) {
	resp, err := s.client.GetItems(ctx, &stockpb.GetItemsRequest{ItemIds: itemsId})
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}
