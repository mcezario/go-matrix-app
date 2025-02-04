package array_utils

import (
	"errors"
	"reflect"
	"testing"
)

type flattenTest struct {
	arg      [][]int
	expected []int
}

type invertTest struct {
	arg, expected [][]int
}

type mathOperationTest struct {
	arg      [][]int
	expected int
}

func Test_Flatten_With_Valid_Matrices(t *testing.T) {
	var flattenTests = []flattenTest{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[][]int{{1, 2}}, []int{1, 2}},
	}

	for _, test := range flattenTests {
		output := Flatten(test.arg)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_Invert_With_Valid_Matrices(t *testing.T) {
	var invertTests = []invertTest{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2}, {4, 5}}, [][]int{{1, 4}, {2, 5}}},
	}

	for _, test := range invertTests {
		output, err := Invert(test.arg)
		if err != nil {
			t.Errorf("Error not expected for matrix %q", test.arg)
		}
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_Invert_With_Wrong_Matrix_Dimension(t *testing.T) {
	output, err := Invert([][]int{{1, 2}})

	if output != nil {
		t.Errorf("An error was expected, got %q instead", output)
	}

	var arrayErr *ArrayUtilsError
	if !errors.As(err, &arrayErr) {
		t.Errorf("An instance of %T is expected, got %T instead", arrayErr, err)
	}

	if err.Error() != "to perform the invert operation, the matrix needs to have the same number of rows and columns" {
		t.Errorf("An error is expected but output %q and error %q were returned", output, err)
	}
}

func Test_Multiply_With_Valid_Matrices(t *testing.T) {
	var multiplyTests = []mathOperationTest{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 362880},
		{[][]int{{1, 2}, {4, 5}}, 40},
	}

	for _, test := range multiplyTests {
		output := Multiply(test.arg)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func Test_Sum_With_Valid_Matrices(t *testing.T) {
	var sumTests = []mathOperationTest{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 45},
		{[][]int{{1, 2}, {4, 5}}, 12},
	}

	for _, test := range sumTests {
		output := Sum(test.arg)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
