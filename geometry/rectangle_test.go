package geometry_test

import (
	"image"
	"testing"

	"github.com/pires/go-dojo-rectangles/geometry"
)

var (
	rectangleA = geometry.Rectangle{
		Name: "A",
		Inner: image.Rectangle{
			Min: image.Point{X: 1, Y: 2},
			Max: image.Point{X: 2, Y: 4},
		},
	}

	rectangleB = geometry.Rectangle{
		Name: "B",
		Inner: image.Rectangle{
			Min: image.Point{X: 2, Y: 3},
			Max: image.Point{X: 5, Y: 4},
		},
	}

	rectangleC = geometry.Rectangle{
		Name: "C",
		Inner: image.Rectangle{
			Min: image.Point{X: 4, Y: 4},
			Max: image.Point{X: 5, Y: 5},
		},
	}

	rectangleD = geometry.Rectangle{
		Name: "D",
		Inner: image.Rectangle{
			Min: image.Point{X: 2, Y: 0},
			Max: image.Point{X: 4, Y: 1},
		},
	}

	rectangleE = geometry.Rectangle{
		Name: "E",
		Inner: image.Rectangle{
			Min: image.Point{X: 4, Y: 2},
			Max: image.Point{X: 6, Y: 3},
		},
	}

	rectangleF = geometry.Rectangle{
		Name: "F",
		Inner: image.Rectangle{
			Min: image.Point{X: -3, Y: -5},
			Max: image.Point{X: -2, Y: -1},
		},
	}

	rectangleG = geometry.Rectangle{
		Name: "G",
		Inner: image.Rectangle{
			Min: image.Point{X: -4, Y: -4},
			Max: image.Point{X: -1, Y: -3},
		},
	}

	rectangleH = geometry.Rectangle{
		Name: "H",
		Inner: image.Rectangle{
			Min: image.Point{X: 1, Y: -4},
			Max: image.Point{X: 4, Y: -1},
		},
	}

	rectangleI = geometry.Rectangle{
		Name: "I",
		Inner: image.Rectangle{
			Min: image.Point{X: 2, Y: -3},
			Max: image.Point{X: 3, Y: -2},
		},
	}

	rectangleJ = geometry.Rectangle{
		Name: "J",
		Inner: image.Rectangle{
			Min: image.Point{X: -4, Y: -5},
			Max: image.Point{X: 7, Y: 6},
		},
	}

	rectangleK = geometry.Rectangle{
		Name: "K",
		Inner: image.Rectangle{
			Min: image.Point{X: -5, Y: 5},
			Max: image.Point{X: -3, Y: 7},
		},
	}
)

type rectangleTuple struct {
	r1 geometry.Rectangle
	r2 geometry.Rectangle
}

func TestRectangle_Height(t *testing.T) {
	var computations = []struct {
		r      geometry.Rectangle
		result int
	}{
		{
			rectangleA,
			2,
		},
		{
			rectangleB,
			1,
		},
	}

	for _, tt := range computations {
		result := tt.r.Height()
		if result != tt.result {
			t.Fatalf("TestRectangle_Height: expected: %d, found: %d", tt.result, result)
		}
	}
}

func TestRectangle_HeightContains(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			// Rect1 height contains Rect2 height
			rectangleTuple{rectangleA, rectangleB},
			true,
		},
		{
			// But Rect1 height is not contained in Rect2 height
			rectangleTuple{rectangleB, rectangleA},
			false,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.HeightContains(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_HeightContains: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestRectangle_Width(t *testing.T) {
	var computations = []struct {
		r      geometry.Rectangle
		result int
	}{
		{
			rectangleA,
			1,
		},
		{
			rectangleB,
			3,
		},
	}

	for _, tt := range computations {
		result := tt.r.Width()
		if result != tt.result {
			t.Fatalf("TestRectangle_Width: expected: %d, found: %d", tt.result, result)
		}
	}
}

func TestRectangle_WidthContains(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			// Rect1 width contains Rect2 width
			rectangleTuple{rectangleB, rectangleC},
			true,
		},
		{
			// But Rect1 width is not contained in Rect2 width
			rectangleTuple{rectangleC, rectangleB},
			false,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.WidthContains(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_WidthContains: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestRectangle_Contains(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			rectangleTuple{rectangleH, rectangleI},
			true,
		},
		{
			// Reverse rectangle order
			rectangleTuple{rectangleI, rectangleH},
			false,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.Contains(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_Contains: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestRectangle_IsContainedIn(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			rectangleTuple{rectangleH, rectangleI},
			false,
		},
		{
			// Reverse rectangle order
			rectangleTuple{rectangleI, rectangleH},
			true,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.IsContainedIn(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_IsContainedIn: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestRectangle_Intersects(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			rectangleTuple{rectangleH, rectangleI},
			false,
		},
		{
			rectangleTuple{rectangleF, rectangleG},
			true,
		},
		{
			rectangleTuple{rectangleK, rectangleJ},
			true,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.Intersects(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_Intersects: expected: %t, found: %t", tt.result, result)
		}
	}
}

func TestRectangle_IntersectionPoints(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result []geometry.Point
	}{
		{
			rectangleTuple{rectangleK, rectangleJ},
			[]geometry.Point{
				{X: -4, Y: 5},
				{X: -3, Y: 6},
			},
		},
		{
			rectangleTuple{rectangleJ, rectangleK},
			[]geometry.Point{
				{X: -4, Y: 5},
				{X: -3, Y: 6},
			},
		},
		{
			rectangleTuple{rectangleF, rectangleG},
			[]geometry.Point{
				{X: -3, Y: -4},
				{X: -3, Y: -3},
				{X: -2, Y: -4},
				{X: -2, Y: -3},
			},
		},
		{
			rectangleTuple{rectangleG, rectangleF},
			[]geometry.Point{
				{X: -3, Y: -4},
				{X: -3, Y: -3},
				{X: -2, Y: -4},
				{X: -2, Y: -3},
			},
		},
		{
			rectangleTuple{rectangleH, rectangleI},
			[]geometry.Point{},
		},
		{
			rectangleTuple{rectangleA, rectangleB},
			[]geometry.Point{},
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.IntersectionPoints(tt.tuple.r2)
		if len(result) != len(tt.result) {
			t.Fatalf("TestRectangle_IntersectionPoints: expected size: %d, found: %d", len(tt.result), len(result))
		}
		for i := 0; i < len(result); i++ {
			if result[i] != tt.result[i] {
				t.Fatalf("TestRectangle_IntersectionPoints: expected: %+v, found: %+v", tt.result[i], result[i])
			}
		}
	}
}

func TestRectangle_IsAdjacent(t *testing.T) {
	var computations = []struct {
		tuple  rectangleTuple
		result bool
	}{
		{
			rectangleTuple{rectangleA, rectangleB},
			true,
		},
		{
			rectangleTuple{rectangleB, rectangleC},
			true,
		},
		{
			rectangleTuple{rectangleB, rectangleE},
			false,
		},
		{
			rectangleTuple{rectangleJ, rectangleG},
			false,
		},
	}

	for _, tt := range computations {
		result := tt.tuple.r1.IsAdjacent(tt.tuple.r2)
		if result != tt.result {
			t.Fatalf("TestRectangle_IsAdjacent: expected: %t, found: %t", tt.result, result)
		}
	}
}
