package ports

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpc    *grpc.Server
	health  *health.Server
	log     *zap.Logger
	address string
}

// NewServer initialises a server from the provided config.
func NewServer(address string, logger *zap.Logger, options ...grpc.ServerOption) (*Server, error) {
	if address == "" {
		return nil, errors.New("server address cannot be empty")
	}

	s := grpc.NewServer(options...)
	hs := health.NewServer()

	// Register the reflection service
	reflection.Register(s)
	healthpb.RegisterHealthServer(s, hs)

	return &Server{
		grpc:    s,
		health:  hs,
		log:     logger,
		address: address,
	}, nil
}

// Run runs the gRPC server until the provided context is cancelled.
// When that happens it shuts down the service.
func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer lis.Close()

	go s.handleShutdown(ctx)

	if err := s.grpc.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("gRPC server failed to start serving: %w", err)
	}
	return nil
}

func (s *Server) handleShutdown(ctx context.Context) {
	<-ctx.Done()
	s.health.Shutdown()

	serverStopped := make(chan struct{})
	go func() {
		s.grpc.GracefulStop()
		close(serverStopped)
	}()
	t := time.NewTimer(10 * time.Second)
	defer t.Stop()

	select {
	case <-t.C:
		s.grpc.Stop()
		s.log.Info("grpc server force stopped")
	case <-serverStopped:
		t.Stop()
		s.log.Info("grpc server gracefully stopped")
	}
}

// RegisterService sets status for the provided service as serving and
// registers the service on the underlying gRPC server.
func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.health.SetServingStatus(desc.ServiceName, healthpb.HealthCheckResponse_SERVING)
	s.grpc.RegisterService(desc, impl)
}
