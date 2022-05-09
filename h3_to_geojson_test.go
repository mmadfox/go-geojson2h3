package geojson2h3

import (
	"testing"

	"github.com/tidwall/geojson"
	"github.com/tidwall/geojson/geometry"
)

func TestToFeatureCollection(t *testing.T) {
	res := 3
	point := geojson.NewPoint(geometry.Point{
		// longitude
		X: 35.29943548054545,
		// latitude
		Y: 101.876220703125,
	})
	indexes, err := ToH3(res, point)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := 1, len(indexes); want != have {
		t.Fatalf("resolution: %d, have %d, want %d", res, have, want)
	}
	featureCollection, err := ToFeatureCollection(indexes)
	if err != nil {
		panic(err)
	}
	if want, have := 1, len(featureCollection.Base()); want != have {
		t.Fatalf("resolution: %d, have %d, want %d", res, have, want)
	}
}
