package main

import (
	"github.com/ChenGuo505/gorder/common/config"
	"github.com/ChenGuo505/gorder/common/logging"
	"github.com/ChenGuo505/gorder/common/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logging.Init()
	if err := config.NewViperConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("payment.service-name")
	serverType := viper.GetString("payment.service-protocol")
	paymentHandler := NewPaymentHandler()

	switch serverType {
	case "http":
		server.RunHTTPServer(serviceName, paymentHandler.RegisterRoutes)
	case "grpc":
		logrus.Panicf("unsupported server type: %s", serverType)
	default:
		logrus.Panicf("unsupported server type: %s", serverType)
	}
}
