package geojson2h3

import (
	"github.com/tidwall/geojson"
	"github.com/tidwall/geojson/geometry"
	"testing"
)

func TestPointToH3(t *testing.T) {
	point := geojson.NewPoint(geometry.Point{
		// longitude
		X: 35.29943548054545,
		// latitude
		Y: 101.876220703125,
	})
	for res := 0; res <= 15; res++ {
		indexes, err := ToH3(res, point)
		if err != nil {
			t.Fatal(err)
		}
		if want, have := 1, len(indexes); want != have {
			t.Fatalf("resolution: %d, have %d, want %d", res, have, want)
		}
	}
}
