package geo2d

// Shape is a geometry shape.
type Shape struct {
	// Shape position.
	position Position
	// Shape dimension.
	dimension Dimension
	// Shape center position (relative to shape position).
	center Position
}

// SetPosition changes shape position coordinates.
func (shape *Shape) SetPosition(x float64, y float64) {
	shape.position.SetCoordinates(x, y)
}

// SetDimension changes shape dimension.
func (shape *Shape) SetDimension(width float64, height float64) {
	shape.dimension.SetValues(width, height)
}

// SetCenter changes shape center coordinates (relative to shape position).
func (shape *Shape) SetCenter(x float64, y float64) {
	shape.center.SetCoordinates(x, y)
}
