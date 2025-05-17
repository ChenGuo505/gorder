package main

import (
	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
	"github.com/ChenGuo505/gorder/common/server"
	"github.com/ChenGuo505/gorder/stock/ports"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.service-protocol")
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer()
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO: implement http server
	default:
		panic("unsupported server type")
	}
}
