package parsers

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

// ParseMatrixCsv receives a file in CSV format that contains a matrix, parses the content, and create a sliced matrix based on the files content.
// Given this string content inside the file:
//
//	1,2,3
//	4,5,6
//	7,8,9
//
// A multimentional array (2d slice) will be returned:
//
//	         0 | 1 | 2
//	    ---------------
//		0  | 1 | 2 | 3
//		1  | 4 | 5 | 6
//		2  | 7 | 8 | 9
//
// It returns a multimentional array (2d slice)
func ParseMatrixCsv(file io.Reader) ([][]int, error) {
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, &ParserError{Err: err}
	}

	// Create the matrix dynamically based on the csv content
	rows, cols := len(records), len(records[0])
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i, row := range records {
		for j, number := range row {
			n, err := strconv.Atoi(number)
			if err != nil {
				return nil, &ParserError{Err: errors.New("wrong csv content: Only numeric input is allowed")}
			}
			matrix[i][j] = n
		}
	}
	return matrix, nil
}
