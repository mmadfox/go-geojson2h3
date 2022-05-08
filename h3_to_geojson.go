package geojson2h3

import (
	"fmt"
	"strconv"

	"github.com/tidwall/geojson"
	"github.com/tidwall/geojson/geometry"
	"github.com/uber/h3-go/v3"
)

func ToFeatureCollection(indexes []h3.H3Index) (*geojson.FeatureCollection, error) {
	if len(indexes) == 0 {
		return nil, fmt.Errorf("uber h3 indexes are empty")
	}
	features := make([]geojson.Object, 0, len(indexes))
	for _, index := range indexes {
		boundary := h3.ToGeoBoundary(index)
		points := make([]geometry.Point, 0, 6)
		for _, b := range boundary {
			points = append(points, geometry.Point{
				X: b.Longitude,
				Y: b.Latitude,
			})
		}
		points = append(points, geometry.Point{
			X: points[0].X,
			Y: points[0].Y,
		})
		polygon := geojson.NewPolygon(
			geometry.NewPoly(points, nil, &geometry.IndexOptions{
				Kind: geometry.None,
			}))
		feature := geojson.NewFeature(polygon, toH3Props(index))
		features = append(features, feature)
	}
	return geojson.NewFeatureCollection(features), nil
}

func toH3Props(index h3.H3Index) string {
	res := strconv.Itoa(h3.Resolution(index))
	return `{"h3index":"` + h3.ToString(index) + `", "h3resolution": ` + res + `}`
}
