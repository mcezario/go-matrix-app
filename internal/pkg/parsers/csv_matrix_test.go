package parsers

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"testing"
)

func Test_Parse_Csv_To_Matrix_Successfully(t *testing.T) {
	file, err := os.Open("../../../tests/data/matrix.csv")
	if err != nil {
		t.Fatalf("File not found: %v", err)
	}
	output, err := ParseMatrixCsv(bufio.NewReader(file))
	if err != nil {
		t.Errorf("A successful matrix is expected but the error %q was returned", err)
	}

	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}

func Test_Parse_Csv_To_Matrix_With_Wrong_Matrix_Dimension(t *testing.T) {
	file, err := os.Open("../../../tests/data/matrix_with_wrong_dimension.csv")
	if err != nil {
		t.Fatalf("File not found: %v", err)
	}
	output, err := ParseMatrixCsv(bufio.NewReader(file))
	if output != nil {
		t.Errorf("An error is expected but a matrix %q was returned", output)
	}

	AssertThatCorrectErrorIsReturned(t, err)
	var expected string = "record on line 2: wrong number of fields"
	if err.Error() != expected {
		t.Errorf("An error equals to %s is expected, got %s instead", expected, err.Error())
	}
}

func Test_Parse_Csv_To_Matrix_With_Wrong_Matrix_Content(t *testing.T) {
	var files = [2]string{
		"../../../tests/data/matrix_with_wrong_content.csv",
		"../../../tests/data/matrix_with_decimal_content.csv",
	}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			t.Fatalf("File not found: %v", err)
		}
		output, err := ParseMatrixCsv(bufio.NewReader(file))
		if output != nil {
			t.Errorf("An error is expected but a matrix %q was returned", output)
		}

		AssertThatCorrectErrorIsReturned(t, err)
		var expected string = "wrong csv content: Only numeric input is allowed"
		if err.Error() != expected {
			t.Errorf("An error equals to %s is expected, got %s instead", expected, err.Error())
		}
	}
}

func AssertThatCorrectErrorIsReturned(t *testing.T, err error) {
	var parserErr *ParserError
	if !errors.As(err, &parserErr) {
		t.Errorf("An instance of %T is expected, got %T instead", parserErr, err)
	}
}
