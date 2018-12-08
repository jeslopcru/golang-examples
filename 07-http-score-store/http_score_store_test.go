package http_score_store

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpScore(t *testing.T) {
	t.Run("return Paco's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		PlayerServer(response, createNewRequest("Paco"))

		assertResponseBody(t, "20", response.Body.String())

	})

	t.Run("return Manolo's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		PlayerServer(response, createNewRequest("Manolo"))

		assertResponseBody(t, "35", response.Body.String())
	})
}

func assertResponseBody(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func createNewRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
