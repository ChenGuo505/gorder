package ports

import (
	"context"

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
	// TODO: implement this method
	panic("implement me")
}

func (g GRPCServer) CheckIfItemsInStock(ctx context.Context, req *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	// TODO: implement this method
	panic("implement me")
}
