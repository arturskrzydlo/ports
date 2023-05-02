package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/arturskrzydlo/ports/internal/ports/adapters"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type portsServiceSuite struct {
	suite.Suite

	repo    Repository
	service pb.PortServiceServer
}

func TestRepository(t *testing.T) {
	suite.Run(t, &portsServiceSuite{})
}

func (s *portsServiceSuite) SetupSuite() {
	log, err := zap.NewDevelopment()
	s.Require().NoError(err)

	s.repo = adapters.NewInMemoryRepo(log)
	s.service = NewPortsService(log, s.repo)
}

func (s *portsServiceSuite) TestStoringPorts() {
	s.Run("should store new port", func() {
		// given
		portToStore := s.createPbPort()

		// when
		_, err := s.service.CreatePort(context.Background(), &pb.CreatePortRequest{Port: portToStore})

		// then
		s.Require().NoError(err)
		portsResp, err := s.service.GetPorts(context.Background(), new(emptypb.Empty))
		s.Assert().Len(portsResp.Ports, 1)
		s.Assert().Equal(s.createPbPort(), portsResp.Ports[0])
	})

	s.Run("should update already existing port", func() {
		// given
		portToStore := s.createPbPort()
		_, err := s.service.CreatePort(context.Background(), &pb.CreatePortRequest{Port: portToStore})

		updatedPort := s.createPbPort()
		updatedPort.Name = "updated-name"

		// when
		_, err = s.service.CreatePort(context.Background(), &pb.CreatePortRequest{Port: updatedPort})

		// then
		s.Require().NoError(err)
		portsResp, err := s.service.GetPorts(context.Background(), new(emptypb.Empty))
		s.Assert().Len(portsResp.Ports, 1)
		s.Assert().Equal(updatedPort.Name, portsResp.Ports[0].Name)
	})
}

func (s *portsServiceSuite) createPbPort() *pb.Port {
	return &pb.Port{
		Name:        "name",
		City:        "London",
		Country:     "UK",
		Alias:       []string{"alias"},
		Regions:     []string{"regions"},
		Coordinates: []float64{90.0},
		Province:    "province",
		Timezone:    "UTC",
		Unlocs:      []string{"unloc"},
		Code:        "some-code",
		Id:          "some-id",
	}
}
