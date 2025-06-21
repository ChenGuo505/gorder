package ports

import (
	"context"
	"github.com/ChenGuo505/gorder/stock/app/query"

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
	items, err := g.app.Queries.GetItems.Handle(ctx, query.GetItems{ItemIds: req.GetItemIds()})
	if err != nil {
		return nil, err
	}
	return &stockpb.GetItemsResponse{Items: items}, nil
}

func (g GRPCServer) CheckIfItemsInStock(ctx context.Context, req *stockpb.CheckIfItemsInStockRequest) (*stockpb.CheckIfItemsInStockResponse, error) {
	items, err := g.app.Queries.CheckIfItemsInStock.Handle(ctx, query.CheckIfItemsInStock{Items: req.GetItems()})
	if err != nil {
		return nil, err
	}
	return &stockpb.CheckIfItemsInStockResponse{
		InStock: 1,
		Items:   items,
	}, nil
}
