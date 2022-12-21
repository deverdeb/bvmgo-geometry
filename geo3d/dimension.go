package geo3d

// Dimension is a dimension (width, height and depth values).
type Dimension struct {
	// Width
	Width float64
	// Height
	Height float64
	// Depth
	Depth float64
}

// Values returns width and height dimension values.
func (dimension *Dimension) Values() (width float64, height float64, depth float64) {
	return dimension.Width, dimension.Height, dimension.Depth
}

// SetValues changes width and height dimension values.
func (dimension *Dimension) SetValues(width float64, height float64, depth float64) {
	dimension.Width, dimension.Height, dimension.Depth = width, height, depth
}
