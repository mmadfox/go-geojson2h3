package main

import (
	"github.com/mmadfox/go-geojson2h3"
	"github.com/tidwall/geojson"
	"github.com/uber/h3-go/v3"
)

func polygonToH3(res int) ([]h3.H3Index, error) {
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
}`, nil)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}

func multiPolygonToH3(res int) ([]h3.H3Index, error) {
	o, err := geojson.Parse(`{
    "type": "MultiPolygon",
    "coordinates": [
        [
            [
                [
                    107,
                    7
                ],
                [
                    108,
                    7
                ],
                [
                    108,
                    8
                ],
                [
                    107,
                    8
                ],
                [
                    107,
                    7
                ]
            ]
        ],
        [
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
    ]
}`, nil)
	checkError(err)
	return geojson2h3.ToH3(res, o)
}
