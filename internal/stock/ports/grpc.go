package ports

import (
	"context"
	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	"github.com/sirupsen/logrus"

	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
	"github.com/ChenGuo505/gorder/stock/app"
)

type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{app: app}
}

func (g GRPCServer) GetItems(ctx context.Context, req *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	logrus.Info("rpc_request_in, stock.GetItems")
	defer func() {
		logrus.Info("rpc_request_out, stock.GetItems")
	}()
	fake := []*orderpb.Item{
		{Id: "fake_item_id"},
	}
	return &stockpb.GetItemsResponse{Items: fake}, nil
}

func (g GRPCServer) CheckIfItemsInStock(ctx context.Context, req *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	logrus.Info("rpc_request_in, stock.CheckIfItemsInStock")
	defer func() {
		logrus.Info("rpc_request_out, stock.CheckIfItemsInStock")
	}()
	return &stockpb.CheckIfItemsInStockResponse{}, nil
}
