package ports

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/arturskrzydlo/ports/internal/grpc"
)

type apiServer struct {
	grpc.UnimplementedPortServiceServer
	log *zap.Logger
}

func NewEligibilityService(log *zap.Logger) grpc.PortServiceServer {
	return &apiServer{
		log: log,
	}
}

func (s *apiServer) CreatePort(ctx context.Context, req *grpc.CreatePortRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
