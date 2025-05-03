package ports

import (
	"context"
	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
)

type GRPCServer struct{}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (g GRPCServer) GetItems(ctx context.Context, req *stockpb.GetItemsRequest) (*stockpb.GetItemsResponse, error) {
	// TODO: implement this method
	panic("implement me")
}

func (g GRPCServer) CheckIfItemsInStock(ctx context.Context, req *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	// TODO: implement this method
	panic("implement me")
}
