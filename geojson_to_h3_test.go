package geojson2h3

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/tidwall/geojson"
	"github.com/tidwall/geojson/geometry"
	"github.com/uber/h3-go/v3"
)

var debug bool = true

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

func TestMultiPointToH3(t *testing.T) {
	res := 7
	points := []geometry.Point{
		{X: -74.143609, Y: 40.751389},
		{X: -73.923951, Y: 40.547124},
		{X: -73.737928, Y: 40.75451},
		{X: -73.902672, Y: 40.764915},
		{X: -74.10311, Y: 40.603463},
	}
	multiPoint := geojson.NewMultiPoint(points)
	indexes, err := ToH3(res, multiPoint)
	if err != nil {
		t.Fatal(err)
	}
	if debug {
		filename := fmt.Sprintf("tmp/multi_point.s1.res:%d.json", res)
		writeIndexesToFile(t, filename, indexes)
	}
	if want, have := len(points), len(indexes); want != have {
		t.Fatalf("resolution: %d, have %d, want %d", res, have, want)
	}
}

func TestLineStringS1ToH3(t *testing.T) {
	points := strToPoints(`
[-73.992074, 40.719831],
[-73.992026, 40.719949]
`)

	testCases := []struct {
		name string
		want int
		res  int
		err  bool
	}{
		{
			name: "success. resolution 0",
			want: 1,
			res:  0,
		},
		{
			name: "success. resolution 1",
			want: 1,
			res:  1,
		},
		{
			name: "success. resolution 2",
			want: 1,
			res:  2,
		},
		{
			name: "success. resolution 3",
			want: 1,
			res:  3,
		},
		{
			name: "success. resolution 4",
			want: 1,
			res:  4,
		},
		{
			name: "success. resolution 5",
			want: 1,
			res:  5,
		},
		{
			name: "success. resolution 6",
			want: 1,
			res:  6,
		},
		{
			name: "success. resolution 7",
			want: 1,
			res:  7,
		},
		{
			name: "success. resolution 8",
			want: 1,
			res:  8,
		},
		{
			name: "success. resolution 9",
			want: 1,
			res:  9,
		},
		{
			name: "success. resolution 10",
			want: 1,
			res:  10,
		},
		{
			name: "success. resolution 11",
			want: 1,
			res:  11,
		},
		{
			name: "success. resolution 12",
			want: 2,
			res:  12,
		},
		{
			name: "success. resolution 13",
			want: 3,
			res:  13,
		},
		{
			name: "success. resolution 14",
			want: 7,
			res:  14,
		},
		{
			name: "success. resolution 15",
			want: 17,
			res:  15,
		},
		{
			name: "failed. resolution -1",
			res:  -1,
			err:  true,
		},
		{

			name: "failed. resolution 16",
			res:  16,
			err:  true,
		},
	}

	if debug {
		_ = os.Mkdir("tmp", 0755)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lineString := geojson.NewLineString(geometry.NewLine(points, nil))
			indexes, err := ToH3(tc.res, lineString)
			if err != nil {
				if !tc.err {
					t.Fatal(err)
				} else {
					return
				}
			}
			if debug {
				filename := fmt.Sprintf("tmp/line_string.s1.res:%d.json", tc.res)
				writeIndexesToFile(t, filename, indexes)
			}
			if have, want := len(indexes), tc.want; have != want {
				t.Fatalf("have %d, want %d", have, want)
			}
		})
	}
}

func TestLineStringS2ToH3(t *testing.T) {
	points := strToPoints(`
[-74.010794, 40.729827],
[-73.932541, 40.67698],
[-73.914179, 40.735812],
[-73.927221, 40.717725],
[-73.938375, 40.742186],
[-73.937689, 40.725663],
[-73.949015, 40.734771],
[-73.942494, 40.705361],
[-73.955879, 40.716424],
[-73.96017, 40.740885],
[-73.995349, 40.745178]
`)

	testCases := []struct {
		name string
		want int
		res  int
		err  bool
	}{
		{
			name: "success. resolution 0",
			want: 1,
			res:  0,
		},
		{
			name: "success. resolution 1",
			want: 1,
			res:  1,
		},
		{
			name: "success. resolution 2",
			want: 1,
			res:  2,
		},
		{
			name: "success. resolution 3",
			want: 1,
			res:  3,
		},
		{
			name: "success. resolution 4",
			want: 2,
			res:  4,
		},
		{
			name: "success. resolution 5",
			want: 2,
			res:  5,
		},
		{
			name: "success. resolution 6",
			want: 4,
			res:  6,
		},
		{
			name: "success. resolution 7",
			want: 13,
			res:  7,
		},
		{
			name: "success. resolution 8",
			want: 43,
			res:  8,
		},
		{
			name: "success. resolution 9",
			want: 112,
			res:  9,
		},
		{
			name: "success. resolution 10",
			want: 306,
			res:  10,
		},
		{
			name: "success. resolution 11",
			want: 827,
			res:  11,
		},
		{
			name: "success. resolution 12",
			want: 2156,
			res:  12,
		},
		{
			name: "success. resolution 13",
			want: 5756,
			res:  13,
		},
		{
			name: "success. resolution 14",
			want: 15183,
			res:  14,
		},
		{
			name: "success. resolution 15",
			want: 40235,
			res:  15,
		},
		{
			name: "failed. resolution -1",
			res:  -1,
			err:  true,
		},
		{

			name: "failed. resolution 16",
			res:  16,
			err:  true,
		},
	}

	if debug {
		_ = os.Mkdir("tmp", 0755)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lineString := geojson.NewLineString(geometry.NewLine(points, nil))
			indexes, err := ToH3(tc.res, lineString)
			if err != nil {
				if !tc.err {
					t.Fatal(err)
				} else {
					return
				}
			}
			if debug {
				filename := fmt.Sprintf("tmp/line_string.s2.res:%d.json", tc.res)
				writeIndexesToFile(t, filename, indexes)
			}
			if have, want := len(indexes), tc.want; have != want {
				t.Fatalf("have %d, want %d", have, want)
			}
		})
	}
}

func writeIndexesToFile(t *testing.T, filename string, indexes []h3.H3Index) {
	featureCollection, err := ToFeatureCollection(indexes)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(featureCollection.JSON()), 0755)
	if err != nil {
		t.Fatal(err)
	}
}

func strToPoints(str string) []geometry.Point {
	lines := strings.Split(str, "\n")
	points := make([]geometry.Point, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line = strings.TrimFunc(line, func(r rune) bool {
			return r == 91 || r == 93 || r == 44
		})
		digits := strings.Split(line, ",")
		if len(digits) == 0 {
			continue
		}
		lon, err := strconv.ParseFloat(strings.TrimSpace(digits[0]), 64)
		if err != nil {
			panic(err)
		}
		lat, err := strconv.ParseFloat(strings.TrimSpace(digits[1]), 64)
		if err != nil {
			panic(err)
		}
		points = append(points, geometry.Point{
			X: lon,
			Y: lat,
		})
	}
	return points
}
