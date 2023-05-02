package webapp

import (
	"encoding/json"
	"fmt"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type Port struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func decodePort(decoder *json.Decoder) (*Port, error) {
	token, decoderErr := decoder.Token()
	if decoderErr != nil {
		return nil, fmt.Errorf("failed to get token: %w", decoderErr)
	}

	var port *Port

	switch typedToken := token.(type) {
	case json.Delim:
		// Do nothing for delimiters like "[" and "]"
	case string: // json should start with port id key which is string
		decodeErr := decoder.Decode(&port)
		if decodeErr != nil {
			return nil, fmt.Errorf("failed to decode to port: %w", decodeErr)
		}
		port.ID = typedToken
	default:
		return nil, fmt.Errorf("incorrect json token: %v", token)
	}

	return port, nil
}

func portToPB(port *Port) *pb.Port {
	return &pb.Port{
		Id:          port.ID,
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
	}
}

func pbToPort(portPb *pb.Port) *Port {
	return &Port{
		ID:          portPb.Id,
		Name:        portPb.Name,
		City:        portPb.City,
		Country:     portPb.Country,
		Alias:       portPb.Alias,
		Regions:     portPb.Regions,
		Coordinates: portPb.Coordinates,
		Province:    portPb.Province,
		Timezone:    portPb.Timezone,
		Unlocs:      portPb.Unlocs,
		Code:        portPb.Code,
	}
}
