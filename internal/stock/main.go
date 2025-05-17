package main

import (
	"context"

	"github.com/ChenGuo505/gorder/common/config"
	"github.com/ChenGuo505/gorder/common/genproto/stockpb"
	"github.com/ChenGuo505/gorder/common/server"
	"github.com/ChenGuo505/gorder/stock/ports"
	"github.com/ChenGuo505/gorder/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.service-protocol")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application := service.NewApplication(ctx)
	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO: implement http server
	default:
		panic("unsupported server type")
	}
}
