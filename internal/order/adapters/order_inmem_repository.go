package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	domain "github.com/ChenGuo505/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	s := make([]*domain.Order, 0)
	s = append(s, &domain.Order{
		Id:          "fakeId",
		CustomerId:  "fakeCustomerId",
		Status:      "fakeStatus",
		PaymentLink: "fakePaymentLink",
		Items:       nil,
	})
	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

func (m *MemoryOrderRepository) Create(_ context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	order.Id = strconv.FormatInt(time.Now().Unix(), 10)
	newOrder := &domain.Order{
		Id:          order.Id,
		CustomerId:  order.CustomerId,
		Status:      order.Status,
		PaymentLink: order.PaymentLink,
		Items:       order.Items,
	}

	m.store = append(m.store, newOrder)

	logrus.WithFields(logrus.Fields{
		"input_order":        order,
		"store_after_create": m.store,
	}).Debug("Order created successfully")

	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(_ context.Context, id string, customerId string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, order := range m.store {
		if order.Id == id && order.CustomerId == customerId {
			logrus.Debugf("Order found: %v", order)
			return order, nil
		}
	}

	return nil, &domain.NotFoundError{Id: id}
}

func (m *MemoryOrderRepository) Update(ctx context.Context, order *domain.Order, updateFunc func(context.Context, *domain.Order) (*domain.Order, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i, o := range m.store {
		if o.Id == order.Id {
			updatedOrder, err := updateFunc(ctx, order)
			if err != nil {
				return err
			}
			m.store[i] = updatedOrder

			logrus.Debugf("Order updated: %v", updatedOrder)

			return nil
		}
	}

	return &domain.NotFoundError{Id: order.Id}
}
