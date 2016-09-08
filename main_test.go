package main

import (
	"bytes"
	"image"
	"testing"

	"github.com/pires/go-dojo-rectangles/geometry"
)

func TestParseJsonFile(t *testing.T) {
	var jsonFiles = []struct {
		filepath      string
		expectedError error
	}{
		{
			"testdata/valid.json",
			nil,
		},
		{
			"testdata/non_existent.json",
			errCantReadFile,
		},
		{
			"testdata/invalid.json",
			errCantDecodeJSON,
		},
		{
			"testdata/not_enough_rectangles.json",
			errNotEnoughRectangles,
		},
	}

	for _, tt := range jsonFiles {
		if _, err := loadRectanglesFromJSONFile(tt.filepath); err != tt.expectedError {
			t.Fatalf("TestParseJsonFile: expected %s, actual %s", tt.expectedError, err)
		}
	}
}

func TestParseValidJson(t *testing.T) {
	var jsonRectangles = []struct {
		json               []byte
		expectedRectangles map[string]geometry.Rectangle
	}{
		{
			[]byte(`[
				{
				  "name": "Rect1",
				  "p1": {
				    "x": 0,
				    "y": 0
				  },
				  "p2": {
				    "x": 10,
				    "y": 10
				  }
				},
				{
				  "name": "Rect2",
				  "p1": {
				    "x": 0,
				    "y": 0
				  },
				  "p2": {
				    "x": 20,
				    "y": 20
				  }
				}
			]`),
			map[string]geometry.Rectangle{
				"Rect1": {
					Name: "Rect1",
					Inner: image.Rectangle{
						Min: image.Point{X: 0, Y: 0},
						Max: image.Point{X: 10, Y: 10},
					},
				},
				"Rect2": {
					Name: "Rect2",
					Inner: image.Rectangle{
						Min: image.Point{X: 0, Y: 0},
						Max: image.Point{X: 20, Y: 20},
					},
				},
			},
		},
	}

	for _, tt := range jsonRectangles {
		if loadedRectangles, err := loadRectanglesFromJSON(bytes.NewReader(tt.json)); err != nil {
			t.Fatalf("TestParseValidJson: unexpected error: %s", err)
		} else {
			// Assert we have loaded the expected number of rectangles
			if len(loadedRectangles) != len(tt.expectedRectangles) {
				t.Fatalf("TestParseValidJson: Expected: %v rectangles but got: %v", len(tt.expectedRectangles), len(loadedRectangles))
			}

			// Iterate loadedRectangles and compare to expected map[string]rectangles
			for name, loadedRectangle := range loadedRectangles {
				if expectedRectangle, ok := tt.expectedRectangles[name]; !ok {
					t.Fatalf("TestParseValidJson: Found unexpected rectangle: %s", name)
				} else {
					if !loadedRectangle.Inner.Eq(expectedRectangle.Inner) {
						t.Fatalf("TestParseValidJson: Loaded rectangle %#+v doesn't match expected rectangle %#+v", loadedRectangle, expectedRectangle)
					}
				}
			}

		}

	}
}
