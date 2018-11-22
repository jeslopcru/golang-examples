package http_score

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpScore(t *testing.T) {
	t.Run("return Paco's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Paco", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		expected := "20"
		result := response.Body.String()

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}
