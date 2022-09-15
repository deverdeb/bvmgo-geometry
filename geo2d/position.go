package geo2d

// Position is a 2D position - (x,y) coordinates.
type Position struct {
	// X is the x coordinate (horizontal).
	X float64
	// Y is the y coordinate (vertical).
	Y float64
}

// Coordinates returns x and y position coordinates.
func (position *Position) Coordinates() (x float64, y float64) {
	return position.X, position.Y
}

// SetCoordinates changes x and y position coordinates.
func (position *Position) SetCoordinates(x float64, y float64) {
	position.X, position.Y = x, y
}
