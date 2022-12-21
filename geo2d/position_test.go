package geo2d

import (
	"reflect"
	"testing"
)

func TestPosition_Coordinates(t *testing.T) {
	position := Position{10, 15}
	x, y := position.Coordinates()
	if x != position.X {
		t.Errorf("position.Coordinates() : x = %v, want %v", x, position.X)
	}
	if y != position.Y {
		t.Errorf("position.Coordinates() : y = %v, want %v", y, position.Y)
	}
}

func TestPosition_SetCoordinates(t *testing.T) {
	position := Position{10, 15}
	position.SetCoordinates(5, 7)
	want := Position{5, 7}
	if !reflect.DeepEqual(position, want) {
		t.Errorf("position.SetCoordinates() = %v, want %v", position, want)
	}
}
