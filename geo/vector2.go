package geo

import (
	"math"
)

// Vec2 is a vector with 2 values: x ,y.
type Vec2 struct {
	// X is the x coordinate (horizontal).
	X float64
	// Y is the y coordinate (vertical).
	Y float64
}

// Add returns a new vector. New vector is the addition of vectors coordinates.
func (vector Vec2) Add(otherVector Vec2) Vec2 {
	return Vec2{X: vector.X + otherVector.X, Y: vector.Y + otherVector.Y}
}

// Sub returns a new vector. New vector is the subtraction of vectors coordinates.
func (vector Vec2) Sub(otherVector Vec2) Vec2 {
	return Vec2{X: vector.X - otherVector.X, Y: vector.Y - otherVector.Y}
}

// MulByVec2 returns a new vector. New vector is the multiplication of vectors coordinates.
func (vector Vec2) MulByVec2(otherVector Vec2) Vec2 {
	return Vec2{X: vector.X * otherVector.X, Y: vector.Y * otherVector.Y}
}

// DivByVec2 returns a new vector. New vector is the division of vectors coordinates.
func (vector Vec2) DivByVec2(otherVector Vec2) Vec2 {
	return Vec2{X: vector.X / otherVector.X, Y: vector.Y / otherVector.Y}
}

// Mul returns a new vector. New vector is the multiplication of vector coordinates and a number.
func (vector Vec2) Mul(number float64) Vec2 {
	return Vec2{X: vector.X * number, Y: vector.Y * number}
}

// Div returns a new vector. New vector is the division of vector coordinates and a number.
func (vector Vec2) Div(number float64) Vec2 {
	return Vec2{X: vector.X / number, Y: vector.Y / number}
}

// Dot returns dot product: v1.v2 = |v1||v2|cos(theta).
func (vector Vec2) Dot(otherVector Vec2) float64 {
	return vector.X*otherVector.X + vector.Y*otherVector.Y
}

// LenSqr returns vector length at square: |v|².
func (vector Vec2) LenSqr() float64 {
	return vector.X*vector.X + vector.Y*vector.Y
}

// Len returns vector length: |v|.
func (vector Vec2) Len() float64 {
	return math.Sqrt(vector.LenSqr())
}

// Normalize returns a new vector with length equals 1 and same vector direction.
func (vector Vec2) Normalize() Vec2 {
	return vector.Mul(1. / vector.Len())
}
