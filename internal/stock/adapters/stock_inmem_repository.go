package adapters

import (
	"context"
	"sync"

	"github.com/ChenGuo505/gorder/common/genproto/orderpb"
	domain "github.com/ChenGuo505/gorder/stock/domain/stock"
)

type MemoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*orderpb.Item
}

var stub = map[string]*orderpb.Item{
	"item_id": {
		Id:       "item_id",
		Name:     "item_name",
		Quantity: 100,
		PriceId:  "price_id",
	},
}

func NewMemoryStockRepository() *MemoryStockRepository {
	return &MemoryStockRepository{
		lock:  &sync.RWMutex{},
		store: stub,
	}
}

func (m *MemoryStockRepository) GetItems(_ context.Context, ids []string) ([]*orderpb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var res []*orderpb.Item
	var missingIds []string

	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missingIds = append(missingIds, id)
		}
	}

	if len(res) == len(ids) {
		return res, nil
	}

	return res, &domain.NotFoundError{Ids: missingIds}
}
