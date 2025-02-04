package api

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type EndpointTest struct {
	endpoint    string
	handlerFunc func(w http.ResponseWriter, r *http.Request)
}

var endpointTests = map[string]EndpointTest{
	"echo":     {"/echo", Echo},
	"invert":   {"/invert", Invert},
	"flatten":  {"/flatten", Flatten},
	"sum":      {"/sum", Sum},
	"multiply": {"/multiply", Multiply},
}

const compliantCsvFile = "../../tests/data/matrix.csv"
const nonCompliantCsvFile = "../../tests/data/matrix_with_wrong_dimension.csv"

func setupTest(t *testing.T, csvFile string, endpoint string, hf func(w http.ResponseWriter, r *http.Request), sendFile bool) (func(tb testing.TB), *httptest.ResponseRecorder) {
	var req *http.Request
	if sendFile {
		file, err := os.Open(csvFile)
		if err != nil {
			t.Fatalf("File not found: %v", err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", csvFile)
		if err != nil {
			t.Fatalf("Failed to create form file: %v", err)
		}

		// Copy the file content into the form field
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatalf("Failed to write CSV file to form: %v", err)
		}

		writer.Close()

		req = httptest.NewRequest("POST", endpoint, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req = httptest.NewRequest("POST", endpoint, nil)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(hf)
	h.ServeHTTP(rr, req)

	return func(tb testing.TB) {
		log.Println("Teardown test")
	}, rr
}

func Test_Echo_Endpoint(t *testing.T) {
	_, rr := BuildSuccessfulScenarioAndMakeRequest(t, compliantCsvFile, "echo")

	AssertThatHttpResponseStatusIsOK(t, rr.Code)

	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func Test_Invert_Endpoint(t *testing.T) {
	_, rr := BuildSuccessfulScenarioAndMakeRequest(t, compliantCsvFile, "invert")

	AssertThatHttpResponseStatusIsOK(t, rr.Code)

	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func Test_Flatten_Endpoint(t *testing.T) {
	_, rr := BuildSuccessfulScenarioAndMakeRequest(t, compliantCsvFile, "flatten")

	AssertThatHttpResponseStatusIsOK(t, rr.Code)

	expected := "1,2,3,4,5,6,7,8,9"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func Test_Sum_Endpoint(t *testing.T) {
	_, rr := BuildSuccessfulScenarioAndMakeRequest(t, compliantCsvFile, "sum")

	AssertThatHttpResponseStatusIsOK(t, rr.Code)

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func Test_Multiply_Endpoint(t *testing.T) {
	_, rr := BuildSuccessfulScenarioAndMakeRequest(t, compliantCsvFile, "multiply")

	AssertThatHttpResponseStatusIsOK(t, rr.Code)

	expected := "362880"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func Test_Endpoints_With_Wrong_Csv_Content(t *testing.T) {
	for _, test := range endpointTests {
		_, rr := BuildFailureScenarioAndMakeRequest(t, nonCompliantCsvFile, test)

		AssertThatHttpResponseStatusIsBadRequest(t, rr.Code)
		expected := "error record on line 2: wrong number of fields"
		if rr.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
		}
	}
}

func Test_Endpoints_With_Csv_Present_In_Request_Body(t *testing.T) {
	for _, test := range endpointTests {
		_, rr := BuildFailureScenarioAndMakeRequest(t, "", test)

		AssertThatHttpResponseStatusIsBadRequest(t, rr.Code)
		expected := "error request Content-Type isn't multipart/form-data"
		if rr.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got %q want %q", rr.Body.String(), expected)
		}
	}
}

func BuildSuccessfulScenarioAndMakeRequest(t *testing.T, csvFile string, testCaseScenario string) (func(tb testing.TB), *httptest.ResponseRecorder) {
	test := endpointTests[testCaseScenario]
	return setupTest(t, csvFile, test.endpoint, test.handlerFunc, true)
}

func BuildFailureScenarioAndMakeRequest(t *testing.T, csvFile string, test EndpointTest) (func(tb testing.TB), *httptest.ResponseRecorder) {
	var sendFile bool = true
	if csvFile == "" {
		sendFile = false
	}

	return setupTest(t, csvFile, test.endpoint, test.handlerFunc, sendFile)
}

func AssertThatHttpResponseStatusIsOK(t *testing.T, statusCode int) {
	if status := statusCode; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func AssertThatHttpResponseStatusIsBadRequest(t *testing.T, statusCode int) {
	if status := statusCode; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
