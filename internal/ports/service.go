package ports

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type apiServer struct {
	pb.UnimplementedPortServiceServer
	log *zap.Logger
}

func NewEligibilityService(log *zap.Logger) pb.PortServiceServer {
	return &apiServer{
		log: log,
	}
}

func (s *apiServer) CreatePort(ctx context.Context, req *pb.CreatePortRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
