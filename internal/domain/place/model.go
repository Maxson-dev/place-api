package place

import "time"

type Place struct {
	ID int64
	Attributes
}

type Attributes struct {
	Coord
	Name      string
	CreatedAt time.Time
}

type Coord struct {
	Lat float64
	Lng float64
}

func NewAttributes(name string, lat, lng float64) Attributes {
	return Attributes{
		Name: name,
		Coord: Coord{
			Lat: lat,
			Lng: lng,
		},
		CreatedAt: time.Now().UTC(),
	}
}
