package ports

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	domainPort "github.com/arturskrzydlo/ports/internal/ports/domain/port"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type apiServer struct {
	pb.UnimplementedPortServiceServer
	log  *zap.Logger
	repo Repository
}

// this service could be separated out from grpc service to have a service layer separate
// from api layer
func NewPortsService(log *zap.Logger, repo Repository) pb.PortServiceServer {
	return &apiServer{
		log:  log,
		repo: repo,
	}
}

func (s *apiServer) CreatePort(ctx context.Context, req *pb.CreatePortRequest) (*emptypb.Empty, error) {
	s.log.Debug("creating port", zap.Any("port", req.Port))
	port, err := portPBToPort(req.Port)
	if err != nil {
		return new(emptypb.Empty), fmt.Errorf("failed to create port:%w", err)
	}

	err = s.repo.CreatePort(ctx, port)
	if err != nil {
		return nil, fmt.Errorf("failed to store port: %w", err)
	}
	return new(emptypb.Empty), nil
}

func (s *apiServer) GetPorts(ctx context.Context, empty *emptypb.Empty) (*pb.GetPortsResponse, error) {
	s.log.Debug("fetching list of ports")
	ports, err := s.repo.GetPorts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all ports: %w", err)
	}
	return portToResponsePayload(ports), nil
}

func portPBToPort(pbPort *pb.Port) (*domainPort.Port, error) {
	return domainPort.NewPort(
		pbPort.Id,
		pbPort.Name,
		pbPort.City,
		pbPort.Country,
		pbPort.Alias,
		pbPort.Regions,
		pbPort.Coordinates,
		pbPort.Province,
		pbPort.Timezone,
		pbPort.Unlocs,
		pbPort.Code)
}

func portToResponsePayload(ports []*domainPort.Port) *pb.GetPortsResponse {
	pbPorts := make([]*pb.Port, len(ports))
	for i, port := range ports {
		pbPorts[i] = &pb.Port{
			Name:        port.Name,
			City:        port.City,
			Country:     port.Country,
			Alias:       port.Alias,
			Regions:     port.Regions,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Timezone:    port.Timezone,
			Unlocs:      port.Unlocs,
			Code:        port.Code,
			Id:          port.ID,
		}
	}
	return &pb.GetPortsResponse{Ports: pbPorts}
}
