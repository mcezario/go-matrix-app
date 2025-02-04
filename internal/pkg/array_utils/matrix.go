package array_utils

import "errors"

// Flatten converts a given matrix in a array. Example:
// This given matrix:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// Will be transformed in:
//
//	1,2,3,4,5,6,7,8,9
//
// It returns a matrix in its flatten format
func Flatten(matrix [][]int) []int {
	var arrayIndex int = 0
	var arraySize int = len(matrix) * len(matrix[0])
	flattenMatrix := make([]int, arraySize)
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			flattenMatrix[arrayIndex] = matrix[i][j]
			arrayIndex++
		}
	}
	return flattenMatrix
}

// Invert inverts the numbers of a given matrix - from left to right, columns become rows and rows become columns. Example:
// This given matrix:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// Becomes:
//
//	1,4,7
//	2,5,8
//	3,6,9
//
// It returns the inverted matrix or an error if the inputed matrix contains diferent sizes of rows and columns.
func Invert(matrix [][]int) ([][]int, error) {
	var matrixSize int = len(matrix)
	if matrixSize != len(matrix[0]) {
		return nil, &ArrayUtilsError{Err: errors.New("to perform the invert operation, the matrix needs to have the same number of rows and columns")}
	}

	invertedMatrix := make([][]int, matrixSize)
	for i := range len(matrix) {
		invertedMatrix[i] = make([]int, matrixSize)
		for j := range len(matrix[i]) {
			invertedMatrix[i][j] = matrix[j][i]
		}
	}
	return invertedMatrix, nil
}

// Invert inverts the numbers of a given matrix - from left to right, columns become rows and rows become columns. Example:
// This given matrix:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// Returns: 362880
//
// It returns multiplication of all numers withi the given matrix.
func Multiply(matrix [][]int) int {
	var multiply int = 1
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			multiply *= matrix[i][j]
		}
	}
	return multiply
}

// Sum sums all numbers of a given matrix. Example:
// This given matrix:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// Returns: 45
//
// It returns the sum of all values within the given matrix.
func Sum(matrix [][]int) int {
	var sum int = 0
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			sum += matrix[i][j]
		}
	}
	return sum
}
