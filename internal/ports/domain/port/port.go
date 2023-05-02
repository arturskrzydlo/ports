package port

type Port struct {
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

func NewPort(name string,
	city string,
	country string,
	alias []string,
	regions []string,
	coordinates []float64,
	province string,
	timezone string,
	unlocs []string,
	code string,
) *Port {
	return &Port{
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
	}
}
