package geo3d

// Position is a 3D position - (x,y,z) coordinates.
type Position struct {
	// X is the x coordinate (horizontal).
	X float64
	// Y is the y coordinate (vertical).
	Y float64
	// Z is the z coordinate (depth).
	Z float64
}

// Coordinates returns x and y position coordinates.
func (position *Position) Coordinates() (x float64, y float64, z float64) {
	return position.X, position.Y, position.Z
}

// SetCoordinates changes x and y position coordinates.
func (position *Position) SetCoordinates(x float64, y float64, z float64) {
	position.X, position.Y, position.Z = x, y, z
}
