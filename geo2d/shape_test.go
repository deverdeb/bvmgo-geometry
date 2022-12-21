package geo2d

import (
	"reflect"
	"testing"
)

func TestBuildShapeFromVertices(t *testing.T) {
	vertices := []Position{
		{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 0, Y: 1},
	}
	shape := BuildShapeFromVertices(vertices)

	if !reflect.DeepEqual(shape.Size(), 4) {
		t.Errorf("shape.Size() = %v, want %v", shape.Size(), 4)
		return
	}
	want1 := Position{0, 0}
	if !reflect.DeepEqual(shape.Vertices[0], want1) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Vertices[0], want1)
	}
	want4 := Position{0, 1}
	if !reflect.DeepEqual(shape.Vertices[3], want4) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Vertices[3], want4)
	}
	boxTL := Position{0, 0}
	if !reflect.DeepEqual(shape.TopLeft, boxTL) {
		t.Errorf("shape.TopLeft = %v, want %v", shape.Vertices[0], boxTL)
	}
	boxBR := Position{1, 1}
	if !reflect.DeepEqual(shape.RightBottom, boxBR) {
		t.Errorf("shape.RightBottom = %v, want %v", shape.Vertices[4], boxBR)
	}
	pos := Position{0., 0.}
	if !reflect.DeepEqual(shape.Position, pos) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Position, pos)
	}
}

func TestBuildShape(t *testing.T) {
	shape := BuildShape(5)

	if !reflect.DeepEqual(shape.Size(), 5) {
		t.Errorf("shape.Size() = %v, want %v", shape.Size(), 5)
		return
	}
	want := Position{0, 0}
	if !reflect.DeepEqual(shape.Vertices[0], want) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Vertices[0], want)
	}
	if !reflect.DeepEqual(shape.Vertices[4], want) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Vertices[4], want)
	}
	if !reflect.DeepEqual(shape.TopLeft, want) {
		t.Errorf("shape.TopLeft = %v, want %v", shape.Vertices[0], want)
	}
	if !reflect.DeepEqual(shape.RightBottom, want) {
		t.Errorf("shape.RightBottom = %v, want %v", shape.Vertices[4], want)
	}
	if !reflect.DeepEqual(shape.Position, want) {
		t.Errorf("position.SetCoordinates() = %v, want %v", shape.Position, want)
	}
}
