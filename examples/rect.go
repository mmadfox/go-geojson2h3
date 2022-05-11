package main

import (
	"github.com/mmadfox/go-geojson2h3"
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func rectToH3(res int) ([]h3.H3Index, error) {
	opts := &geojson.ParseOptions{
		AllowRects: true,
	}
	o, err := geojson.Parse(`{
    "type": "Polygon",
    "coordinates": [
        [
            [
                100,
                0
            ],
            [
                101,
                0
            ],
            [
                101,
                1
            ],
            [
                100,
                1
            ],
            [
                100,
                0
            ]
        ]
    ]
}`, opts)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}
