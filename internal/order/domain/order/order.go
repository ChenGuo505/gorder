package order

import "github.com/ChenGuo505/gorder/common/genproto/orderpb"

type Order struct {
	Id          string
	CustomerId  string
	Status      string
	PaymentLink string
	Items       []*orderpb.Item
}
