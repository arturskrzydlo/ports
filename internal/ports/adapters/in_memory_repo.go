package adapters

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/arturskrzydlo/ports/internal/ports/domain/port"
)

type InMemoryRepo struct {
	sync.RWMutex
	log     *zap.Logger
	storage map[string]*port.Port
}

func NewInMemoryRepo(logger *zap.Logger) *InMemoryRepo {
	return &InMemoryRepo{
		log:     logger,
		storage: make(map[string]*port.Port),
	}
}

func (r *InMemoryRepo) CreatePort(ctx context.Context, port *port.Port) error {
	r.Lock()
	defer r.Unlock()
	r.storage[port.ID] = port
	return nil
}

func (r *InMemoryRepo) GetPorts(ctx context.Context) ([]*port.Port, error) {
	r.RLock()
	defer r.RUnlock()
	ports := make([]*port.Port, 0)
	for _, storagePort := range r.storage {
		if storagePort != nil {
			ports = append(ports, storagePort)
		}
	}
	return ports, nil
}
