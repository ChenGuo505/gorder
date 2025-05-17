package order

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(context.Context, *Order) (*Order, error)
	Get(context.Context, string, string) (*Order, error)
	Update(context.Context, *Order, func(context.Context, *Order) (*Order, error)) error
}

type NotFoundError struct {
	Id string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("order with ID %s not found", e.Id)
}
