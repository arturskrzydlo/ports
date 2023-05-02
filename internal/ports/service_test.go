package ports

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type portsServiceSuite struct {
	suite.Suite

	repo    PortsRepo
	service pb.PortServiceServer
}

func TestRepository(t *testing.T) {
	suite.Run(t, &portsServiceSuite{})
}

func (s *portsServiceSuite) SetupSuite() {
	log, err := zap.NewDevelopment()
	s.Require().NoError(err)
	s.service = NewPortsService(log)
}

func (s *portsServiceSuite) TestStoringPorts() {
	// given
	portToStore := pb.Port{
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

	// when
	_, err := s.service.CreatePort(context.Background(), &pb.CreatePortRequest{Port: &portToStore})

	// then
	ports, err := s.repo.GetPorts(context.Background())
	s.Assert().Len(ports, 1)
	s.Assert().Equal(portToStore, ports[0])
}
