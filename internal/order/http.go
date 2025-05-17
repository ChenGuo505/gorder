package main

import (
	"net/http"

	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	"github.com/ChenGuo505/gorder/order/app"
	"github.com/ChenGuo505/gorder/order/app/command"
	"github.com/ChenGuo505/gorder/order/app/query"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {
	var req orderpb.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	r, err := s.app.Commands.CreateOrder.Handle(c, command.CreateOrder{
		CustomerId: req.CustomerId,
		Items:      req.Items,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": r})
}

func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	order, err := s.app.Queries.GetCustomerOrder.Handle(c, query.GetCustomerOrder{
		CustomerId: customerId,
		OrderId:    orderId,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}
