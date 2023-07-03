package http_request

import (
	"net/http"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	url := "https://www.google.com"
	responseBody, statusCode := HttpRequest("GET", url, nil)

	if statusCode != http.StatusOK {
		t.Fatalf("Invalid status code: %d", statusCode)
	}

	actualBody := string(responseBody)

	if len(actualBody) == 0 {
		t.Fatalf("Invalid body: %d", statusCode)
	}
}
