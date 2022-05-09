package main

import (
	"fmt"

	"github.com/mmadfox/go-geojson2h3"
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func main() {
	resolution := 9
	object, err := geojson.Parse(`{"type":"FeatureCollection","features":[{"type":"Feature","properties":{"shape":"Polygon","name":"Unnamed Layer","category":"default"},"geometry":{"type":"Polygon","coordinates":[[[-73.901303,40.756892],[-73.893924,40.743755],[-73.871476,40.756278],[-73.863378,40.764175],[-73.871444,40.768467],[-73.879852,40.760014],[-73.885515,40.764045],[-73.891522,40.761054],[-73.901303,40.756892]]]},"id":"a6ca1b7e-9ddf-4425-ad07-8a895f7d6ccf"}]}`, nil)
	if err != nil {
		panic(err)
	}

	indexes, err := geojson2h3.ToH3(resolution, object)
	if err != nil {
		panic(err)
	}
	for _, index := range indexes {
		fmt.Printf("h3index: %s\n", h3.ToString(index))
	}

	featureCollection, err := geojson2h3.ToFeatureCollection(indexes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Polyfill:")
	fmt.Println(featureCollection.JSON())
}
