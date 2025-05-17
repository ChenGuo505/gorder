package stock

import (
	"context"

	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
)

type Repository interface {
	GetItems(context.Context, []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Ids []string
}

func (e *NotFoundError) Error() string {
	return "items not found"
}
