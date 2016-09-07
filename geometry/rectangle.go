package geometry

import (
	"image"
)

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Rectangle struct {
	Name string
	P1   point `json: p1`
	P2   point `json: p2`

	Inner image.Rectangle `json:"-"`
}

// Height returns this rectangle height.
func (this *Rectangle) Height() int {
	return this.Inner.Max.Y - this.Inner.Min.Y
}

// HeightContains whether this rectangle height contains that rectangle height in the Y axis
func (this *Rectangle) HeightContains(that Rectangle) bool {
	return this.Inner.Min.Y <= that.Inner.Min.Y && this.Inner.Max.Y >= that.Inner.Max.Y
}

// Width returns this rectangle width.
func (this *Rectangle) Width() int {
	return this.Inner.Max.X - this.Inner.Min.X
}

// WidthContains whether this rectangle contains tht rectangle width in the X axis
func (this *Rectangle) WidthContains(that Rectangle) bool {
	return this.Inner.Min.X <= that.Inner.Min.X && this.Inner.Max.X >= that.Inner.Max.X
}

// Contains whether this rectangle contains that rectangle.
func (this *Rectangle) Contains(that Rectangle) bool {
	return that.Inner.In(this.Inner)
}

// IsContainedIn whether this rectangle is contained within that rectangle.
func (this *Rectangle) IsContainedIn(that Rectangle) bool {
	return this.Inner.In(that.Inner)
}

// Intersects whether this rectangle intersects that rectangle.
func (this *Rectangle) Intersects(that Rectangle) bool {
	return this.Inner.Overlaps(that.Inner) && (!this.Contains(that) && !this.IsContainedIn(that))
}

// IsAdjacent whether this rectangle is adjacent with that rectangle.
func (this *Rectangle) IsAdjacent(that Rectangle) (result bool) {
	// Reject if one contains the other
	if this.Contains(that) || this.IsContainedIn(that) {
		result = false
	} else {
		// Are rectangles adjacent in the Y axis?
		heightContained := this.HeightContains(that) || that.HeightContains(*this)
		if heightContained {
			// Axis X points (left or right side) from both rectangles must touch
			result = this.Inner.Min.X == that.Inner.Min.X || this.Inner.Min.X == that.Inner.Max.X || this.Inner.Max.X == that.Inner.Min.X || this.Inner.Max.X == that.Inner.Max.X
		}

		widthContained := this.WidthContains(that) || that.WidthContains(*this)
		if widthContained {
			// Axis Y points (upper or lower side) from both rectangles must touch
			result = this.Inner.Min.Y == that.Inner.Min.Y || this.Inner.Min.Y == that.Inner.Max.Y || this.Inner.Max.Y == that.Inner.Min.Y || this.Inner.Max.Y == that.Inner.Max.Y
		}
	}

	return
}
