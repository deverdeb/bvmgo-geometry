package geo

import (
	"math"
	"testing"
)

func TestTransformationMat3x3Builder_Identity(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	if !testEqualVec3(result, vec) {
		t.Errorf("Identity transformation = %v, want %v", result, vec)
	}
}

func TestTransformationMat3x3Builder_Rotation_45(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(45.).Build()
	vec := Vec3{2., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{0., math.Sqrt(8), 1.} // math.Sqrt(8) = vector size = sqrt(2²+2²)
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Rotation_Minus45(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(-45.).Build()
	vec := Vec3{2., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{math.Sqrt(8), 0, 1.} // math.Sqrt(8) = vector size = sqrt(2²+2²)
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Rotation_90(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(90.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{-2., 1., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Rotation_180(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(180.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{-1., -2., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Rotation_270(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(270.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{2., -1., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Rotation_360(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(360.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{1., 2., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Rotation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Scale(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Scale(2., 3.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{2., 6., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Scale transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Scale_0(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Scale(0., 0.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{0., 0., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Scale transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Scale_1(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Scale(1., 1.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{1., 2., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Scale transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_Translation(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Translation(Vec2{2., 3.}).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{3., 5., 1.}
	if !testEqualVec3(result, want) {
		t.Errorf("Translation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_ScaleAndTranslation(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Scale(2, 3).Translation(Vec2{-1, -2}).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{1., 4., 1.} // (1, 2) --[zoom(2, 3)]-> (2, 6) --[trans(-1, -2)]--> (1, 4)
	if !testEqualVec3(result, want) {
		t.Errorf("Scale and translation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_RotationAndTranslation(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Rotation(90.).Translation(Vec2{-1, -2}).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{-3., -1., 1.} // (1, 2) --[rot(90°)]-> (-2, 1) --[trans(-1, -2)]--> (-3, -1)
	if !testEqualVec3(result, want) {
		t.Errorf("Scale and translation transformation = %v, want %v", result, want)
	}
}

func TestTransformationMat3x3Builder_TranslationAndRotation(t *testing.T) {
	matrix := NewTransformationMat3x3Builder().Translation(Vec2{2, 1}).Rotation(-45.).Build()
	vec := Vec3{1., 2., 1.}
	result := matrix.MulVec3(vec)
	want := Vec3{math.Sqrt(18), 0., 1.} // (1, 2) --[trans(2, 1)]--> (3, 3) --[rot(-45°)]-> (sqrt(3²+3²), 0)
	if !testEqualVec3(result, want) {
		t.Errorf("Scale and translation transformation = %v, want %v", result, want)
	}
}

// testEqualVec3 returns true if vectrices are equal (with threshold to compare float64).
func testEqualVec3(v1, v2 Vec3) bool {
	const equalityThreshold = 1e-9
	eq := func(a, b float64) bool {
		return math.Abs(a-b) <= equalityThreshold
	}
	return eq(v1.X, v2.X) && eq(v1.Y, v2.Y) && eq(v1.Z, v2.Z)
}
