package webhook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIError represents discord answer to the request if it went wrong.
// https://discord.com/developers/docs/topics/opcodes-and-status-codes
type APIError struct {
	HTTPResponse int    `json:"http_response"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
}

// Error implements `error` interface.
func (e *APIError) Error() string {
	return e.Message
}

// CheckResponse is a help func to check responses from discord API for errors
// `body` contains response from the server and probably will contain error information if
// it hasn't been parsed to `*APIError` type.
func CheckResponse(response *http.Response) (body []byte, err error) {
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = response.Body.Close()
	if err != nil {
		return
	}
	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNoContent:
	case http.StatusCreated:
		return
	default:
		result := &APIError{
			HTTPResponse: response.StatusCode,
		}
		err = json.Unmarshal(body, result)
		if err != nil {
			return
		}
		err = result
	}
	return
}

func CheckError(response *http.Response) error {
	_, err := CheckResponse(response)
	return err
}