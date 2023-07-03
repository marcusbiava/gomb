package http_request

import (
	"net/http"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	url := "https://www.google.com"
	responseBody, statusCode, err := HttpRequest("GET", url, nil)

	if err != nil {
		t.Fatalf("Error is not nil %s", err)
	}

	if statusCode != http.StatusOK {
		t.Fatalf("Invalid status code: %d", statusCode)
	}

	actualBody := string(responseBody)

	if len(actualBody) == 0 {
		t.Fatalf("Invalid body: %d", statusCode)
	}
}
