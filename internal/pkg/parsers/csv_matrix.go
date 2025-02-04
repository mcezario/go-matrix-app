package parsers

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

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
