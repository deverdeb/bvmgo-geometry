package geo2d

// Dimension is a dimension (width and height values).
type Dimension struct {
	// Width
	Width float64
	// Height
	Height float64
}

// Values returns width and height dimension values.
func (dimension *Dimension) Values() (width float64, height float64) {
	return dimension.Width, dimension.Height
}

// SetValues changes width and height dimension values.
func (dimension *Dimension) SetValues(width float64, height float64) {
	dimension.Width, dimension.Height = width, height
}
