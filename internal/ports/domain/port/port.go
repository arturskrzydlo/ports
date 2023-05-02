package port

import "errors"

type Port struct {
	ID          string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

func NewPort(id string,
	name string,
	city string,
	country string,
	alias []string,
	regions []string,
	coordinates []float64,
	province string,
	timezone string,
	unlocs []string,
	code string,
) (*Port, error) {
	if id == "" {
		return nil, errors.New("port id can't be empty")
	}
	if code == "" {
		return nil, errors.New("port code can't be empty")
	}
	return &Port{
		ID:          id,
		Name:        name,
		City:        city,
		Country:     country,
		Alias:       alias,
		Regions:     regions,
		Coordinates: coordinates,
		Province:    province,
		Timezone:    timezone,
		Unlocs:      unlocs,
		Code:        code,
	}, nil
}
