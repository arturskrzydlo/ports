package ports

import (
	"context"

	"github.com/arturskrzydlo/ports/internal/ports/domain/port"
)

type Repository interface {
	CreatePort(ctx context.Context, port *port.Port) error
	GetPorts(ctx context.Context) ([]*port.Port, error)
}
