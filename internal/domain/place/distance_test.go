package place

import (
	"testing"
)

func Test_Distance(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name  string
		p     Coord
		q     Coord
		outMi float64
		outKm float64
	}{
		{
			"Rio de Janeiro, Brazil -> Bangkok, Thailand",
			Coord{Lat: 22.55, Lng: 43.12},
			Coord{Lat: 13.45, Lng: 100.28},
			3786.251258825624,
			6094.544408786774,
		},
		{
			"Port Louis, Mauritius -> Padang, Indonesia",
			Coord{Lat: 20.10, Lng: 57.30},
			Coord{Lat: 0.57, Lng: 100.21},
			3196.671009759937,
			5145.525771394785,
		},
		{
			"Oxford, United Kingdom -> Vatican City, Vatican City",
			Coord{Lat: 51.45, Lng: 1.15},
			Coord{Lat: 41.54, Lng: 12.27},
			863.0311907424888,
			1389.1793118293067,
		},
		{
			"Windhoek, Namibia -> Rotterdam, Netherlands",
			Coord{Lat: 22.34, Lng: 17.05},
			Coord{Lat: 51.56, Lng: 4.29},
			2130.8298370015464,
			3429.89310043882,
		},
		{
			"Esperanza, Argentina -> Luanda, Angola",
			Coord{Lat: 63.24, Lng: 56.59},
			Coord{Lat: 8.50, Lng: 13.14},
			4346.398369403186,
			6996.18595539861,
		},
		{
			"North/South Poles -> Paris, France",
			Coord{Lat: 90.00, Lng: 0.00},
			Coord{Lat: 48.51, Lng: 2.21},
			2866.1346681303867,
			4613.477506482742,
		},
		{
			"Turin, Italy -> Kuala Lumpur, Malaysia",
			Coord{Lat: 45.04, Lng: 7.42},
			Coord{Lat: 3.09, Lng: 101.42},
			6261.05275709582,
			10078.111954385415,
		},
		{
			"Russia, Moscow -> Russia, Saint Petersburg",
			Coord{Lat: 55.45, Lng: 37.36},
			Coord{Lat: 59.57, Lng: 30.19},
			389.1861714998077,
			626.4540420983514,
		},
	}

	for _, input := range tests {
		input := input
		t.Run(input.name, func(t *testing.T) {
			t.Parallel()

			mi, km := Distance(input.p, input.q)

			if input.outMi != mi || input.outKm != km {
				t.Errorf("fail: want %v %v -> %v %v got %v %v",
					input.p,
					input.q,
					input.outMi,
					input.outKm,
					mi,
					km,
				)
			}
		})
	}
}
