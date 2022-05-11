package main

import (
	"github.com/mmadfox/go-geojson2h3"
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func circleToH3(res int) ([]h3.H3Index, error) {
	o, err := geojson.Parse(`{
        "type": "Feature",
        "properties": {
            "type": "Circle",
            "radius": 1616.8125889500864,
            "radius_units":"m",
            "name": "Unnamed Layer",
            "category": "default"
        },
        "geometry": {
            "type": "Point",
            "coordinates": [-73.900445, 40.766256]
        },
        "id": "c7823150-b827-4281-b288-51ce333a776d"
    }`, nil)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}
