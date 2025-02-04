package api

import (
	"fmt"
	"mcezario/backend-challenge/internal/pkg/array_utils"
	"mcezario/backend-challenge/internal/pkg/parsers"
	"net/http"
	"strconv"
	"strings"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	matrix := GetMatrixFromCsv(w, r)
	if matrix == nil {
		return
	}
	fmt.Fprint(w, ConvertMatrixToString(matrix))
}

func Invert(w http.ResponseWriter, r *http.Request) {
	matrix := GetMatrixFromCsv(w, r)
	if matrix == nil {
		return
	}

	invertedMatrix, err := array_utils.Invert(matrix)
	if invertedMatrix != nil {
		fmt.Fprint(w, ConvertMatrixToString(invertedMatrix))
	} else {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
	}
}

func Flatten(w http.ResponseWriter, r *http.Request) {
	matrix := GetMatrixFromCsv(w, r)
	if matrix == nil {
		return
	}

	flattenMatrix := array_utils.Flatten(matrix)
	fmt.Fprint(w, strings.Trim(strings.Join(strings.Split(fmt.Sprint(flattenMatrix), " "), ","), "[]"))
}

func Sum(w http.ResponseWriter, r *http.Request) {
	matrix := GetMatrixFromCsv(w, r)
	if matrix == nil {
		return
	}

	var sum int = array_utils.Sum(matrix)
	fmt.Fprint(w, sum)
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	matrix := GetMatrixFromCsv(w, r)
	if matrix == nil {
		return
	}

	var multiply int = array_utils.Multiply(matrix)
	fmt.Fprint(w, strconv.Itoa(multiply))
}

func GetMatrixFromCsv(w http.ResponseWriter, r *http.Request) [][]int {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteErrorToResponse(w, err)
		return nil
	}
	defer file.Close()
	matrix, err := parsers.ParseMatrixCsv(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		WriteErrorToResponse(w, err)
		return nil
	}
	return matrix
}

func ConvertMatrixToString(matrix [][]int) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Trim(strings.Join(strings.Split(fmt.Sprint(row), " "), ","), "[]"))
	}
	return response
}

func WriteErrorToResponse(w http.ResponseWriter, err error) {
	w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
}
