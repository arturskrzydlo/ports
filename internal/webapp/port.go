package webapp

import (
	"encoding/json"
	"fmt"
)

type Port struct {
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

	switch token.(type) {
	case json.Delim:
		// Do nothing for delimiters like "[" and "]"
	case string: // json should start with port name which is string
		decodeErr := decoder.Decode(&port)
		if decodeErr != nil {
			return nil, fmt.Errorf("failed to decode to port: %w", decodeErr)
		}
	default:
		return nil, fmt.Errorf("incorrect json token: %v", token)
	}

	return port, nil
}