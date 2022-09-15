package geo

import "math"

// Vec3 is a vector with 3 values: x ,y, z.
type Vec3 struct {
	// X is the x coordinate (horizontal).
	X float64
	// Y is the y coordinate (vertical).
	Y float64
	// Z is the z coordinate (depth).
	Z float64
}

// Add returns a new vector. New vector is the addition of vectors coordinates.
func (vector Vec3) Add(otherVector Vec3) Vec3 {
	return Vec3{X: vector.X + otherVector.X, Y: vector.Y + otherVector.Y, Z: vector.Z + otherVector.Z}
}

// Sub returns a new vector. New vector is the subtraction of vectors coordinates.
func (vector Vec3) Sub(otherVector Vec3) Vec3 {
	return Vec3{X: vector.X - otherVector.X, Y: vector.Y - otherVector.Y, Z: vector.Z - otherVector.Z}
}

// DivByVec3 returns a new vector. New vector is the division of vectors coordinates.
func (vector Vec3) DivByVec3(otherVector Vec3) Vec3 {
	return Vec3{X: vector.X / otherVector.X, Y: vector.Y / otherVector.Y, Z: vector.Z / otherVector.Z}
}

// Mul returns a new vector. New vector is the multiplication of vector coordinates and a number.
func (vector Vec3) Mul(number float64) Vec3 {
	return Vec3{X: vector.X * number, Y: vector.Y * number, Z: vector.Z * number}
}

// MulByVec3 returns a new vector. New vector is the multiplication of vectors coordinates.
func (vector Vec3) MulByVec3(otherVector Vec3) Vec3 {
	return Vec3{X: vector.X * otherVector.X, Y: vector.Y * otherVector.Y, Z: vector.Z * otherVector.Z}
}

// MulMat3x3 returns a new vector, result of multiplication between vector (matrix 1x3) and matrix.
func (vector Vec3) MulMat3x3(matrix Mat3x3) Vec3 {
	return Vec3{
		matrix[0][0]*vector.X + matrix[1][0]*vector.Y + matrix[2][0]*vector.Z,
		matrix[0][1]*vector.X + matrix[1][1]*vector.Y + matrix[2][1]*vector.Z,
		matrix[0][2]*vector.X + matrix[1][2]*vector.Y + matrix[2][2]*vector.Z,
	}
}

// Div returns a new vector. New vector is the division of vector coordinates and a number.
func (vector Vec3) Div(number float64) Vec3 {
	return Vec3{X: vector.X / number, Y: vector.Y / number, Z: vector.Z / number}
}

// Dot returns dot product: v1.v2 = |v1||v2|cos(theta).
func (vector Vec3) Dot(otherVector Vec3) float64 {
	return vector.X*otherVector.X + vector.Y*otherVector.Y + vector.Z*otherVector.Z // FIXME...
}

// LenSqr returns vector length at square: |v|².
func (vector Vec3) LenSqr() float64 {
	return vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
}

// Len returns vector length: |v|.
func (vector Vec3) Len() float64 {
	return math.Sqrt(vector.LenSqr())
}

// Normalize returns a new vector with length equals 1 and same vector direction.
func (vector Vec3) Normalize() Vec3 {
	return vector.Mul(1. / vector.Len())
}
