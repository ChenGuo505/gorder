package main

import (
	"github.com/ChenGuo505/gorder/order/app"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct{
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {
	// TODO implement me
}

func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	// TODO implement me
}
