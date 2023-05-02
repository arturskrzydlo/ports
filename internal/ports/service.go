package ports

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/arturskrzydlo/ports/internal/common/pb"

	domainPort "github.com/arturskrzydlo/ports/internal/ports/domain/port"
)

type APIServer struct {
	pb2.UnimplementedPortServiceServer
	log  *zap.Logger
	repo Repository
}

// TODO: this service could be separated out from grpc service to have a service layer separate
// from api layer - however it would be really thin in this case
func NewPortsService(log *zap.Logger, repo Repository) *APIServer {
	return &APIServer{
		log:  log,
		repo: repo,
	}
}

func (s *APIServer) CreatePort(ctx context.Context, req *pb2.CreatePortRequest) (*emptypb.Empty, error) {
	s.log.Debug("creating port", zap.Any("port", req.Port))
	port, err := portPBToPort(req.Port)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("failed to create port:%w", err)
	}

	err = s.repo.CreatePort(ctx, port)
	if err != nil {
		return nil, fmt.Errorf("failed to store port: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *APIServer) GetPorts(ctx context.Context, _ *emptypb.Empty) (*pb2.GetPortsResponse, error) {
	s.log.Debug("fetching list of ports")
	ports, err := s.repo.GetPorts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all ports: %w", err)
	}
	return portToResponsePayload(ports), nil
}

func portPBToPort(pbPort *pb2.Port) (*domainPort.Port, error) {
	port, err := domainPort.NewPort(
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
	if err != nil {
		return nil, fmt.Errorf("failed to map pb port to domain port: %w", err)
	}
	return port, nil
}

func portToResponsePayload(ports []*domainPort.Port) *pb2.GetPortsResponse {
	pbPorts := make([]*pb2.Port, len(ports))
	for i, port := range ports {
		pbPorts[i] = &pb2.Port{
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
	return &pb2.GetPortsResponse{Ports: pbPorts}
}
