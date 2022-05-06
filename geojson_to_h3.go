package geojson2h3

import (
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func ToH3(resolution int, o geojson.Object) ([]h3.H3Index, error) {
	switch typ := o.(type) {
	case *geojson.Point:
		return pointToH3(resolution, typ), nil
	}
	return nil, nil
}

func pointToH3(resolution int, point *geojson.Point) []h3.H3Index {
	index := h3.FromGeo(h3.GeoCoord{
		Latitude:  point.Center().Y,
		Longitude: point.Center().X,
	}, resolution)
	return []h3.H3Index{index}
}
