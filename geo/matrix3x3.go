package geo

import "log"

// Mat3x3 is a 3x3 matrix (3 rows x 3 columns).
// First index: row (start at 0) / second index: column (start at 0)
//
//	Mat3x3[3][3]float64 {
//	   {r0c0, r0c1, r0c2}, // row 0
//	   {r1c0, r1c1, r1c2}, // row 1
//	   {r2c0, r2c1, r2c2}, // row 2
//	}
type Mat3x3 [3][3]float64

// Mat3x3Identity returns 3x3 identity matrix.
func Mat3x3Identity() Mat3x3 {
	return Mat3x3{
		{1., 0., 0.}, // row 0
		{0., 1., 0.}, // row 1
		{0., 0., 1.}, // row 2
	}
}

// Mat3x3Zero returns 3x3 matrix with only 0 values.
func Mat3x3Zero() Mat3x3 {
	return Mat3x3{
		{0., 0., 0.}, // row 0
		{0., 0., 0.}, // row 1
		{0., 0., 0.}, // row 2
	}
}

// Add ajoute les valeurs de deux matrices et retourne le résultat sous forme de matrice
func (matrix Mat3x3) Add(otherMat Mat3x3) Mat3x3 {
	return Mat3x3{
		{matrix[0][0] + otherMat[0][0], matrix[0][1] + otherMat[0][1], matrix[0][2] + otherMat[0][2]},
		{matrix[1][0] + otherMat[1][0], matrix[1][1] + otherMat[1][1], matrix[1][2] + otherMat[1][2]},
		{matrix[2][0] + otherMat[2][0], matrix[2][1] + otherMat[2][1], matrix[2][2] + otherMat[2][2]},
	}
}

// Sub soustrait les valeurs de deux matrices et retourne le résultat sous forme de matrice
func (matrix Mat3x3) Sub(otherMat Mat3x3) Mat3x3 {
	return Mat3x3{
		{matrix[0][0] - otherMat[0][0], matrix[0][1] - otherMat[0][1], matrix[0][2] - otherMat[0][2]},
		{matrix[1][0] - otherMat[1][0], matrix[1][1] - otherMat[1][1], matrix[1][2] - otherMat[1][2]},
		{matrix[2][0] - otherMat[2][0], matrix[2][1] - otherMat[2][1], matrix[2][2] - otherMat[2][2]},
	}
}

// Mul multiplie les valeurs de d'une matrice par une valeur et retourne le résultat sous forme de matrice
func (matrix Mat3x3) Mul(value float64) Mat3x3 {
	return Mat3x3{
		{matrix[0][0] * value, matrix[0][1] * value, matrix[0][2] * value},
		{matrix[1][0] * value, matrix[1][1] * value, matrix[1][2] * value},
		{matrix[2][0] * value, matrix[2][1] * value, matrix[2][2] * value},
	}
}

// MulVec3 returns a new vector, result of multiplication between matrix and vector (matrix 3x1).
func (matrix Mat3x3) MulVec3(vector Vec3) Vec3 {
	return Vec3{
		matrix[0][0]*vector.X + matrix[0][1]*vector.Y + matrix[0][2]*vector.Z,
		matrix[1][0]*vector.X + matrix[1][1]*vector.Y + matrix[1][2]*vector.Z,
		matrix[2][0]*vector.X + matrix[2][1]*vector.Y + matrix[2][2]*vector.Z,
	}
}

// MulMat3x3 returns a new matrix, result of multiplication between matrices.
func (matrix Mat3x3) MulMat3x3(otherMatrix Mat3x3) Mat3x3 {
	return Mat3x3{
		{
			matrix[0][0]*otherMatrix[0][0] + matrix[0][1]*otherMatrix[1][0] + matrix[0][2]*otherMatrix[2][0],
			matrix[0][0]*otherMatrix[0][1] + matrix[0][1]*otherMatrix[1][1] + matrix[0][2]*otherMatrix[2][1],
			matrix[0][0]*otherMatrix[0][2] + matrix[0][1]*otherMatrix[1][2] + matrix[0][2]*otherMatrix[2][2],
		}, {
			matrix[1][0]*otherMatrix[0][0] + matrix[1][1]*otherMatrix[1][0] + matrix[1][2]*otherMatrix[2][0],
			matrix[1][0]*otherMatrix[0][1] + matrix[1][1]*otherMatrix[1][1] + matrix[1][2]*otherMatrix[2][1],
			matrix[1][0]*otherMatrix[0][2] + matrix[1][1]*otherMatrix[1][2] + matrix[1][2]*otherMatrix[2][2],
		}, {
			matrix[2][0]*otherMatrix[0][0] + matrix[2][1]*otherMatrix[1][0] + matrix[2][2]*otherMatrix[2][0],
			matrix[2][0]*otherMatrix[0][1] + matrix[2][1]*otherMatrix[1][1] + matrix[2][2]*otherMatrix[2][1],
			matrix[2][0]*otherMatrix[0][2] + matrix[2][1]*otherMatrix[1][2] + matrix[2][2]*otherMatrix[2][2],
		},
	}
}

// Row method returns matrix row (array of 3 values).
// First index is 0, max index is 2. Invalid error value throws fatal error log with os.Exit(1).
func (matrix Mat3x3) Row(index int) [3]float64 {
	if index < 0 || index > 2 {
		log.Fatalf("Mat3x3.Row(index) : index %d out of bound, required number between 0 and 2 (include)", index)
	}
	return [3]float64{matrix[index][0], matrix[index][1], matrix[index][2]}
}

// Column method returns matrix column (array of 3 values).
// First index is 0, max index is 2. Invalid error value throws fatal error log with os.Exit(1).
func (matrix Mat3x3) Column(index int) [3]float64 {
	if index < 0 || index > 2 {
		log.Fatalf("Mat3x3.Column(index) : index %d out of bound, required number between 0 and 2 (include)", index)
	}
	return [3]float64{matrix[0][index], matrix[1][index], matrix[0][index]}
}
