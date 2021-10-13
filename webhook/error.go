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
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = response.Body.Close()
	if err != nil {
		return
	}
	body = resp
	if response.StatusCode != http.StatusOK &&
		response.StatusCode != http.StatusNoContent &&
		response.StatusCode != http.StatusCreated {
		result := &APIError{
			HTTPResponse: response.StatusCode,
		}
		err = json.Unmarshal(resp, result)
		if err != nil {
			return
		}
		err = result
		return
	}
	return
}
