package webhook

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	err := &APIError{
		Message: "Hello world",
	}
	if err.Error() != err.Message {
		t.Errorf("Expected %s, got %s", err.Message, err.Error())
	}
}

func TestCheckResponse(t *testing.T) {
	testData := map[string]struct {
		in  *http.Response
		out *APIError
	}{
		"Code 200 test": {
			in: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(&bytes.Buffer{}),
			},
			out: nil,
		},
		"Code 400 test": {
			in: &http.Response{
				StatusCode: 400,
				Body: ioutil.NopCloser(
					bytes.NewBufferString(`{"code": 50006, "message": "Cannot send an empty message"}`)),
			},
			out: &APIError{
				HTTPResponse: 400,
				Code:         50006,
				Message:      "Cannot send an empty message",
			},
		},
	}
	for test, td := range testData {
		t.Run(test, func(t *testing.T) {
			_, err := CheckResponse(td.in)
			if td.out == nil {
				if err != nil {
					t.Error("Want nil, got ", err)
				}
			} else {
				if apiError, ok := err.(*APIError); ok {
					switch {
					case apiError.HTTPResponse != td.out.HTTPResponse:
						t.Errorf("Expected %d response, got %d", td.out.HTTPResponse, apiError.HTTPResponse)
					case apiError.Code != td.out.Code:
						t.Errorf("Expected %d code, got %d", td.out.Code, apiError.Code)
					case apiError.Message != td.out.Message:
						t.Errorf("Expected %s, got %s", td.out.Message, apiError.Message)
					}
				} else {
					t.Error("Got wrong error")
				}
			}
		})
	}
}
