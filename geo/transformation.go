package geo

import "math"

// TransformationMat3x3Builder build a transformation matrix.
type TransformationMat3x3Builder struct {
	matrix Mat3x3
}

// NewTransformationMat3x3Builder returns a new TransformationMat3x3Builder.
func NewTransformationMat3x3Builder() *TransformationMat3x3Builder {
	return &TransformationMat3x3Builder{
		matrix: Mat3x3Identity(),
	}
}

// Translation adds translation effect to transformation matrix.
func (builder *TransformationMat3x3Builder) Translation(vec Vec2) *TransformationMat3x3Builder {
	matrixTranslation := Mat3x3{
		{1, 0, vec.X},
		{0, 1, vec.Y},
		{0, 0, 1},
	}
	builder.matrix = matrixTranslation.MulMat3x3(builder.matrix)
	return builder
}

// Rotation adds rotation effect to transformation matrix.
func (builder *TransformationMat3x3Builder) Rotation(angleInDeg float64) *TransformationMat3x3Builder {
	angleInDeg360 := math.Mod(angleInDeg, 360)
	if angleInDeg360 < 0 {
		angleInDeg360 += 360
	}
	var matrixRotation Mat3x3
	if angleInDeg360 != 0 {
		if angleInDeg360 == 90 {
			matrixRotation = Mat3x3{
				{0, -1, 0},
				{1, 0, 0},
				{0, 0, 1},
			}
		} else if angleInDeg360 == 180 {
			matrixRotation = Mat3x3{
				{-1, 0, 0},
				{0, -1, 0},
				{0, 0, 1},
			}
		} else if angleInDeg360 == 270 {
			matrixRotation = Mat3x3{
				{0, 1, 0},
				{-1, 0, 0},
				{0, 0, 1},
			}
		} else {
			angleInRad := DegToRad(angleInDeg360)
			cosinus := math.Cos(angleInRad)
			sinus := math.Sin(angleInRad)
			matrixRotation = Mat3x3{
				{cosinus, -sinus, 0},
				{sinus, cosinus, 0},
				{0, 0, 1},
			}
		}
		builder.matrix = matrixRotation.MulMat3x3(builder.matrix)
	}
	return builder
}

// Scale adds zoom effect to transformation matrix.
func (builder *TransformationMat3x3Builder) Scale(x float64, y float64) *TransformationMat3x3Builder {
	matrixScale := Mat3x3{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1},
	}
	builder.matrix = matrixScale.MulMat3x3(builder.matrix)
	return builder
}

// Build returns the transformation matrix.
func (builder *TransformationMat3x3Builder) Build() Mat3x3 {
	return builder.matrix
}
