package geo2d

import (
	"reflect"
	"testing"
)

func TestDimension_Values(t *testing.T) {
	dimension := Dimension{10, 15}
	w, h := dimension.Values()
	if w != dimension.Width {
		t.Errorf("dimension.Values() : w = %v, want %v", w, dimension.Width)
	}
	if h != dimension.Height {
		t.Errorf("dimension.Values() : h = %v, want %v", h, dimension.Height)
	}
}

func TestDimension_SetValues(t *testing.T) {
	dimension := Dimension{10, 15}
	dimension.SetValues(5, 7)
	want := Dimension{5, 7}
	if !reflect.DeepEqual(dimension, want) {
		t.Errorf("dimension.SetValues() = %v, want %v", dimension, want)
	}
}
