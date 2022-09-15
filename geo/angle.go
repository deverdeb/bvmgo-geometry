package geo

import "math"

// DegToRad convertit des degrés en radians
func DegToRad(angleValue float64) float64 {
	return math.Mod(angleValue, 360) * math.Pi / 180.
}

// RadToDeg convertit des radians en degrés
func RadToDeg(angleValue float64) float64 {
	return math.Mod(angleValue*180/math.Pi, 360)
}
