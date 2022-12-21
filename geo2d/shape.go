package geo2d

import "github.com/deverdeb/bvmgo-geometry/geo"

// Shape is a shape with a position and multiple vertices.
type Shape struct {
	// Position is the absolute shape center position
	Position Position
	// Vertices positions
	// Vertices are relative to shape centre (0, 0)
	Vertices []Position
	// TopLeft is the top/left bounding box corner (relative)
	TopLeft Position
	// TopLeft is the bottom/right bounding box corner (relative)
	RightBottom Position
}

// BuildShapeFromVertices function builds a shape from array of vertices.
// Vertex positions are relative to shape centre (0, 0).
func BuildShapeFromVertices(vertices []Position) Shape {
	shape := Shape{
		Position: Position{X: 0, Y: 0},
		Vertices: vertices,
	}
	shape.ComputeBoundingBox()
	return shape
}

// BuildShape function builds an empty shape.
func BuildShape(size int) Shape {
	shape := Shape{
		Position: Position{X: 0, Y: 0},
		Vertices: make([]Position, size),
	}
	shape.ComputeBoundingBox()
	return shape
}

// Size returns number of vertices
func (shape *Shape) Size() int {
	if shape.Vertices == nil {
		return 0
	}
	return len(shape.Vertices)
}

// VertexRelative returns a relative vertex position (relative to shape center).
func (shape *Shape) VertexRelative(index int) Position {
	if index < 0 || shape.Size() <= index {
		return Position{X: 0, Y: 0}
	}
	return shape.Vertices[index]
}

// VertexAbsolute returns an absolute vertex position (relative to shape position).
func (shape *Shape) VertexAbsolute(index int) Position {
	relativePosition := shape.VertexRelative(index)
	return Position{
		X: shape.Position.X + relativePosition.X,
		Y: shape.Position.Y + relativePosition.Y,
	}
}

// ComputeBoundingBox function computes shape bounding box (min / max positions).
func (shape *Shape) ComputeBoundingBox() {
	if shape.Size() <= 0 {
		shape.TopLeft = Position{X: 0, Y: 0}
		shape.RightBottom = Position{X: 0, Y: 0}
		return
	}
	shape.TopLeft = shape.Vertices[0]
	shape.RightBottom = shape.Vertices[0]
	shape.ForEachVertices(func(_ *Shape, _ int, vertex Position) {
		if vertex.X < shape.TopLeft.X {
			shape.TopLeft.X = vertex.X
		}
		if vertex.Y < shape.TopLeft.Y {
			shape.TopLeft.Y = vertex.Y
		}
		if vertex.X > shape.RightBottom.X {
			shape.RightBottom.X = vertex.X
		}
		if vertex.Y > shape.RightBottom.Y {
			shape.RightBottom.Y = vertex.Y
		}
	})
}

// TranslateCenter function changes shape center and recomputes shape bounding box.
func (shape *Shape) TranslateCenter(translation geo.Vec2) {
	shape.ForEachVertices(func(_ *Shape, index int, vertex Position) {
		shape.Vertices[index] = vertex.Translate(translation)
	})
	shape.ComputeBoundingBox()
}

// MoveCenterTo function changes shape center and recomputes shape box.
// Centre position is relative to current shape centre (0, 0)
func (shape *Shape) MoveCenterTo(centre Position) {
	shape.TranslateCenter(geo.Vec2{X: centre.X, Y: centre.Y})
}

// ForEachVertices function executes parameter function at all vertices.
func (shape *Shape) ForEachVertices(function func(shape *Shape, index int, vertex Position)) {
	if shape.Vertices == nil {
		return
	}
	for index, vertex := range shape.Vertices {
		function(shape, index, vertex)
	}
}
