package main

import (
	"fmt"

	"github.com/mmadfox/go-geojson2h3"

	"github.com/uber/h3-go/v3"
)

type dataSet struct {
	typ string
	res int
	fn  func(res int) ([]h3.H3Index, error)
}

func main() {
	examples := []dataSet{
		{
			typ: "Point",
			res: 6,
			fn:  pointToH3,
		},
		{
			typ: "MultiPoint",
			res: 6,
			fn:  multiPointToH3,
		},
		{
			typ: "Line",
			res: 4,
			fn:  lineToH3,
		},
		{
			typ: "MultiLine",
			res: 10,
			fn:  multiLineToH3,
		},
		{
			typ: "Polygon",
			res: 4,
			fn:  polygonToH3,
		},
		{
			typ: "MultiPolygon",
			res: 4,
			fn:  multiPolygonToH3,
		},
		{
			typ: "Collection",
			res: 9,
			fn:  collectionToH3,
		},
		{
			typ: "Feature",
			res: 8,
			fn:  featureToH3,
		},
		{
			typ: "FeatureCollection",
			res: 8,
			fn:  featureCollectionToH3,
		},
		{
			typ: "Circle",
			res: 8,
			fn:  circleToH3,
		},
	}
	for _, example := range examples {
		fmt.Printf("GeoJSON: %s\n", example.typ)
		indexes, err := example.fn(example.res)
		checkError(err)
		fmt.Printf("indexes: %d\n", len(indexes))
		featureCollection, err := geojson2h3.ToFeatureCollection(indexes)
		if err != nil {
			panic(err)
		}
		fmt.Println("polyfill:")
		fmt.Println(featureCollection.JSON())
		fmt.Println("")
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
