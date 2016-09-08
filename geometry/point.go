package geometry

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// In reports whether p is in r and p is not part of any of r lines.
func (p Point) ContainedButNotInLine(r Rectangle) bool {
	return p.X > r.Inner.Min.X && p.X < r.Inner.Max.X &&
		p.Y > r.Inner.Min.Y && p.Y < r.Inner.Max.Y
}
