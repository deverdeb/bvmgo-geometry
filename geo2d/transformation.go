package geo2d

import "github.com/deverdeb/bvmgo-geometry/geo"

// Translate function move position.
func (position *Position) Translate(vec geo.Vec2) Position {
	return Position{X: position.X + vec.X, Y: position.Y + vec.Y}
}

// Rotation function applies rotation angle to position from rotation center.
func (position *Position) Rotation(center *Position, angleInDeg float64) Position {
	return position.Transform(center, geo.NewTransformationMat3x3Builder().Rotation(angleInDeg).Build())
}

// Scale function applies zoom / scale to position from scale center.
func (position *Position) Scale(center *Position, factor float64) Position {
	return position.Transform(center, geo.NewTransformationMat3x3Builder().Scale(factor, factor).Build())
}

// Transform function applies transformation matrix to position from transformation center.
func (position *Position) Transform(center *Position, transformMatrix geo.Mat3x3) Position {
	initialVector := geo.Vec3{X: position.X - center.X, Y: position.Y - center.Y, Z: 1.}
	transformedVector := transformMatrix.MulVec3(initialVector)
	return Position{X: transformedVector.X, Y: transformedVector.Y}
}

// Scale function applies zoom / scale to dimension.
func (dimension *Dimension) Scale(factor float64) Dimension {
	return Dimension{Width: dimension.Width * factor, Height: dimension.Height * factor}
}

// Translate function move position.
func (shape *Shape) Translate(vec geo.Vec2) Shape {
	return Shape{
		Position:    shape.Position.Translate(vec),
		Vertices:    shape.Vertices,
		TopLeft:     shape.TopLeft,
		RightBottom: shape.RightBottom,
	}
}

// Rotation function applies rotation angle to shape from shape center.
func (shape *Shape) Rotation(angleInDeg float64) Shape {
	if shape.Size() <= 0 {
		return *shape
	}
	vertices := make([]Position, shape.Size())
	centre := Position{0, 0}
	shape.ForEachVertices(func(shape *Shape, index int, vertex Position) {
		vertices[index] = vertex.Rotation(&centre, angleInDeg)
	})
	newShape := BuildShapeFromVertices(vertices)
	newShape.Position = shape.Position
	return newShape
}

// Scale function applies zoom / scale to shape.
func (shape *Shape) Scale(factor float64) Shape {
	if shape.Size() <= 0 {
		return *shape
	}
	vertices := make([]Position, shape.Size())
	centre := Position{0, 0}
	shape.ForEachVertices(func(shape *Shape, index int, vertex Position) {
		vertices[index] = vertex.Scale(&centre, factor)
	})
	newShape := BuildShapeFromVertices(vertices)
	newShape.Position = shape.Position
	return newShape
}
