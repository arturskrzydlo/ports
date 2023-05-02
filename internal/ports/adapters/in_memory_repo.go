package adapters

import (
	"context"
	"sync"

	"go.uber.org/zap"

	domainPort "github.com/arturskrzydlo/ports/internal/ports/domain/port"
)

type InMemoryRepo struct {
	mutex   sync.RWMutex
	log     *zap.Logger
	storage map[string]*domainPort.Port
}

func NewInMemoryRepo(logger *zap.Logger) *InMemoryRepo {
	return &InMemoryRepo{
		log:     logger,
		storage: make(map[string]*domainPort.Port),
	}
}

func (r *InMemoryRepo) CreatePort(_ context.Context, port *domainPort.Port) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.storage[port.ID] = port
	return nil
}

func (r *InMemoryRepo) GetPorts(_ context.Context) ([]*domainPort.Port, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	ports := make([]*domainPort.Port, 0)
	for _, storagePort := range r.storage {
		if storagePort != nil {
			ports = append(ports, storagePort)
		}
	}
	return ports, nil
}
