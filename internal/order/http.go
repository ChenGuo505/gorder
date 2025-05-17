package main

import (
	"net/http"

	"github.com/ChenGuo505/gorder/order/app"
	"github.com/ChenGuo505/gorder/order/app/query"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	app app.Application
}

func (s HTTPServer) PostCustomerCustomerIdOrders(c *gin.Context, customerId string) {
	// TODO: implement me
}

func (s HTTPServer) GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string) {
	order, err := s.app.Queries.GetCustomerOrder.Handle(c, query.GetCustomerOrder{
		CustomerId: "fakeCustomerId",
		OrderId:    "fakeId",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": order})
}
