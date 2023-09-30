package place

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

func (p Place) DistanceToKm(q Place) float64 {
	_, km := Distance(p.Coord, q.Coord)
	return km
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q Coord) (mi, km float64) {
	lat1 := degreesToRadians(p.Lat)
	Lng1 := degreesToRadians(p.Lng)
	lat2 := degreesToRadians(q.Lat)
	Lng2 := degreesToRadians(q.Lng)

	diffLat := lat2 - lat1
	diffLng := Lng2 - Lng1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLng/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * earthRadiusMi
	km = c * earthRaidusKm

	return mi, km
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}
