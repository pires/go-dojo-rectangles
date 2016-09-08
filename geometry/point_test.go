package geometry_test

import (
	"testing"

	"github.com/pires/go-dojo-rectangles/geometry"
)

func TestPoint_ContainedButNotInLinen(t *testing.T) {
	var computations = []struct {
		point     geometry.Point
		rectangle geometry.Rectangle
		result    bool
	}{
		{
			geometry.Point{X: -4, Y: 5},
			rectangleJ,
			false,
		},
		{
			geometry.Point{X: -3, Y: 5},
			rectangleJ,
			true,
		},
		{
			geometry.Point{X: -4, Y: 6},
			rectangleK,
			true,
		},
		{
			geometry.Point{X: -3, Y: 5},
			rectangleK,
			false,
		},
	}

	for _, tt := range computations {
		result := tt.point.ContainedButNotInLine(tt.rectangle)
		if result != tt.result {
			t.Fatalf("TestPoint_ContainedButNotInLinen: expected: %t, found: %t", tt.result, result)
		}
	}
}
