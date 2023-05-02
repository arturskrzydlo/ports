package port

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPortCreation(t *testing.T) {
	tests := map[string]struct {
		ID   string
		code string
		err  bool
	}{
		"should create port when ID and code fields are present": {
			ID:   "some-id",
			code: "some-code",
			err:  false,
		},
		"shouldn't create port when ID is missing": {
			ID:   "",
			code: "some-code",
			err:  true,
		},
		"shouldn't create port when code is missing": {
			ID:   "some-id",
			code: "",
			err:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// given
			name := "port-name"
			city := "London"
			country := "United Kingdom"
			alias := []string{"alias"}
			regions := []string{"region"}
			coordinates := []float64{90.0, 90.0}
			province := "province"
			timezone := "UTC"
			unlocs := []string{"abc"}

			// when
			port, err := NewPort(tc.ID, name, city, country, alias, regions, coordinates, province, timezone, unlocs, tc.code)
			if tc.err {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, Port{
					ID:          tc.ID,
					Name:        name,
					City:        city,
					Country:     country,
					Alias:       alias,
					Regions:     regions,
					Coordinates: coordinates,
					Province:    province,
					Timezone:    timezone,
					Unlocs:      unlocs,
					Code:        tc.code,
				}, *port)
			}
		})
	}
}
