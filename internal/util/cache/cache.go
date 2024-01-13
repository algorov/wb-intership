package cache

import (
	"l0Service/internal/util/jsonutil"
	"sync"
)

// Cache ...
type Cache struct {
	mx    *sync.RWMutex
	cache map[string]jsonutil.Order
}

// New ...
func New() *Cache {
	return &Cache{
		cache: make(map[string]jsonutil.Order),
	}
}

// AddOrder ...
func (s *Cache) AddOrder(order jsonutil.Order) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.cache[order.OrderUID] = order
}

// AddOrder ...
func (s *Cache) GetOrder(orderUID string) {

}
