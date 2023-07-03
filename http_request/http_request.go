package http_request

import (
	"bytes"
	"fmt"
	"github.com/marcusbiava/gomb/errors"
	"io"
	"net/http"
)

// HttpRequest makes an HTTP request using the specified method, URL, and headers.
// It returns the response body as a byte slice and the response status code.
func HttpRequest(method, url string, headers map[string]string) ([]byte, int) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	errors.IfAnErrorOccursCallsLogFatal(err, fmt.Sprintf("http.NewRequest(%s, %s)", method, url))

	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	var body []byte
	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	resp, err := client.Do(req)
	errors.IfAnErrorOccursCallsLogFatal(err, "client.Do(req)")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		errors.IfAnErrorOccursCallsLogFatal(err, "Body.Close()")
	}(resp.Body)

	responseBody, err := io.ReadAll(resp.Body)
	errors.IfAnErrorOccursCallsLogFatal(err, "io.ReadAll(resp.Body)")

	return responseBody, resp.StatusCode
}
