package geometry

import (
	"image"
)

type Rectangle struct {
	Name string
	P1   Point `json: p1`
	P2   Point `json: p2`

	Inner image.Rectangle `json:"-"`
}

// Height returns this rectangle height.
func (this Rectangle) Height() int {
	return this.Inner.Max.Y - this.Inner.Min.Y
}

// HeightContains whether this rectangle height contains that rectangle height in the Y axis
func (this Rectangle) HeightContains(that Rectangle) bool {
	return this.Inner.Min.Y <= that.Inner.Min.Y && this.Inner.Max.Y >= that.Inner.Max.Y
}

// Width returns this rectangle width.
func (this Rectangle) Width() int {
	return this.Inner.Max.X - this.Inner.Min.X
}

// WidthContains whether this rectangle contains tht rectangle width in the X axis
func (this Rectangle) WidthContains(that Rectangle) bool {
	return this.Inner.Min.X <= that.Inner.Min.X && this.Inner.Max.X >= that.Inner.Max.X
}

// Contains whether this rectangle contains that rectangle.
func (this Rectangle) Contains(that Rectangle) bool {
	return that.Inner.In(this.Inner)
}

// IsContainedIn whether this rectangle is contained within that rectangle.
func (this Rectangle) IsContainedIn(that Rectangle) bool {
	return this.Inner.In(that.Inner)
}

// Intersects whether this rectangle intersects that rectangle.
func (this Rectangle) Intersects(that Rectangle) bool {
	return this.Inner.Overlaps(that.Inner) && (!this.Contains(that) && !this.IsContainedIn(that))
}

// IntersectionPoints returns an array of intersection points.
func (this Rectangle) IntersectionPoints(that Rectangle) (intersectionPoints []Point) {
	// Evaluate intersection
	if this.Intersects(that) {
		// Get intersection result, a new rectangle
		intersectionRectangle := this.Inner.Intersect(that.Inner)
		// If rectangles intersect, the intersectionRectangle is not empty but it doesn't hurt to double check
		if !intersectionRectangle.Eq(image.ZR) {
			// Get intersectionRectangle vertices
			vertices := []Point{
				{X: intersectionRectangle.Min.X, Y: intersectionRectangle.Min.Y},
				{X: intersectionRectangle.Min.X, Y: intersectionRectangle.Max.Y},
				{X: intersectionRectangle.Max.X, Y: intersectionRectangle.Min.Y},
				{X: intersectionRectangle.Max.X, Y: intersectionRectangle.Max.Y},
			}
			// Return only intersection points that share both rectangles lines.
			for _, vertice := range vertices {
				if !(vertice.ContainedButNotInLine(this) || vertice.ContainedButNotInLine(that)) {
					intersectionPoints = append(intersectionPoints, vertice)
				}
			}
		}
	}

	return
}

// IsAdjacent whether this rectangle is adjacent with that rectangle.
func (this Rectangle) IsAdjacent(that Rectangle) (result bool) {
	// Reject if one contains the other
	if this.Contains(that) || this.IsContainedIn(that) {
		result = false
	} else {
		// Are rectangles adjacent in the Y axis?
		heightContained := this.HeightContains(that) || that.HeightContains(this)
		if heightContained {
			// Axis X points (left or right side) from both rectangles must touch
			result = this.Inner.Min.X == that.Inner.Min.X || this.Inner.Min.X == that.Inner.Max.X || this.Inner.Max.X == that.Inner.Min.X || this.Inner.Max.X == that.Inner.Max.X
		}

		widthContained := this.WidthContains(that) || that.WidthContains(this)
		if widthContained {
			// Axis Y points (upper or lower side) from both rectangles must touch
			result = this.Inner.Min.Y == that.Inner.Min.Y || this.Inner.Min.Y == that.Inner.Max.Y || this.Inner.Max.Y == that.Inner.Min.Y || this.Inner.Max.Y == that.Inner.Max.Y
		}
	}

	return
}
