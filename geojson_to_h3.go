package geojson2h3

import (
	"fmt"

	"github.com/tidwall/geojson/geo"
	"github.com/tidwall/geojson/geometry"

	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func ToH3(resolution int, o geojson.Object) (indexes []h3.H3Index, err error) {
	if o == nil {
		return nil, fmt.Errorf("geojson.Object is nil")
	}
	if resolution < 0 || resolution > 15 {
		return nil, fmt.Errorf("got invalid resolution %d. expected from 0 to 15",
			resolution)
	}
	switch typ := o.(type) {
	case *geojson.MultiPoint:
		set := make([][]h3.H3Index, 0)
		typ.ForEach(func(object geojson.Object) bool {
			point, ok := object.(*geojson.Point)
			if !ok {
				return false
			}
			indexes := pointToH3(resolution, point)
			set = append(set, indexes)
			return true
		})
		indexes = deDup(set)
	case *geojson.Point:
		return pointToH3(resolution, typ), nil
	case *geojson.MultiLineString:
		set := make([][]h3.H3Index, 0)
		typ.ForEach(func(geom geojson.Object) bool {
			lineString, ok := geom.(*geojson.LineString)
			if !ok {
				return false
			}
			indexes, err = lineStringToH3(resolution, lineString)
			if err != nil {
				return false
			}
			set = append(set, indexes)
			return true
		})
		indexes = deDup(set)
	case *geojson.LineString:
		indexes, err = lineStringToH3(resolution, typ)
		if err != nil {
			return nil, err
		}
		indexes = deDup([][]h3.H3Index{indexes})
	}
	return
}

func pointToH3(resolution int, point *geojson.Point) []h3.H3Index {
	index := h3.FromGeo(h3.GeoCoord{
		Latitude:  point.Center().Y,
		Longitude: point.Center().X,
	}, resolution)
	return []h3.H3Index{index}
}

func lineStringToH3(resolution int, lineString *geojson.LineString) ([]h3.H3Index, error) {
	if lineString.Base().NumPoints() < 2 {
		return nil, fmt.Errorf("got %d points, expected >= 2 points",
			lineString.Base().NumPoints())
	}
	step := stepForResolution(resolution)
	points := make([]geometry.Point, 0, 2)
	for i := 0; i < lineString.Base().NumSegments(); i++ {
		segment := lineString.Base().SegmentAt(i)
		dist := distanceMeters(segment)
		if dist > step {
			points = append(points, segment.A)
			b := bearing(segment)
			for j := float64(0); j < dist; j += step {
				next := j
				if next+step > dist {
					next = dist
				}
				lat, lon := geo.DestinationPoint(segment.A.Y, segment.A.X, next, b)
				points = append(points, geometry.Point{X: lon, Y: lat})
			}
		} else {
			points = append(points, segment.A)
			points = append(points, segment.B)
		}
	}
	indexes := make([]h3.H3Index, 0, len(points))
	for i := 0; i < len(points); i++ {
		cellID := h3.FromGeo(h3.GeoCoord{
			Latitude:  points[i].Y,
			Longitude: points[i].X}, resolution)
		indexes = append(indexes, cellID)
	}
	return indexes, nil
}

func deDup(indexes [][]h3.H3Index) []h3.H3Index {
	if len(indexes) == 0 {
		return []h3.H3Index{}
	}
	visits := make(map[h3.H3Index]struct{})
	result := make([]h3.H3Index, 0)
	for i := 0; i < len(indexes); i++ {
		set := indexes[i]
		for j := 0; j < len(set); j++ {
			idx := set[j]
			if _, ok := visits[idx]; ok {
				continue
			}
			visits[idx] = struct{}{}
			result = append(result, idx)
		}
	}
	return result
}

func distanceMeters(s geometry.Segment) float64 {
	return geo.DistanceTo(s.A.Y, s.A.X, s.B.Y, s.B.X)
}

func bearing(s geometry.Segment) float64 {
	return geo.BearingTo(s.A.Y, s.A.X, s.B.Y, s.B.X)
}

const (
	level0km  = 1107
	level1km  = 418
	level2km  = 158
	level3km  = 59
	level4km  = 22
	level5km  = 8
	level6km  = 3
	level7km  = 1
	level8km  = 0.46
	level9km  = 0.17
	level10km = 0.06
	level11km = 0.024
	level12km = 0.0094
	level13km = 0.0035
	level14km = 0.0013
	level15km = 0.0005
)

var steps = map[int]float64{
	0:  level0km,
	1:  level1km,
	2:  level2km,
	3:  level3km,
	4:  level4km,
	5:  level5km,
	6:  level6km,
	7:  level7km,
	8:  level8km,
	9:  level9km,
	10: level10km,
	11: level11km,
	12: level12km,
	13: level13km,
	14: level14km,
	15: level15km,
}

func stepForResolution(level int) (meters float64) {
	km, ok := steps[level]
	if !ok {
		km = level7km
	}
	meters = km * 1000
	return
}
