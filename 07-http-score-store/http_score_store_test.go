package http_score_store

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) ObtainPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestHttpScore(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Paco":   20,
			"Manolo": 35,
		},
	}
	server := &PlayerServer{&store}

	t.Run("return Paco's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Paco")
		server.ServeHTTP(response, request)

		assertResponseBody(t, "20", response.Body.String())

	})

	t.Run("return Manolo's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Manolo")
		server.ServeHTTP(response, request)

		assertResponseBody(t, "35", response.Body.String())
	})

	t.Run("return 404 when not found a player", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Juan")
		server.ServeHTTP(response, request)

		expected := http.StatusNotFound
		result := response.Code
		if result != expected {
			t.Errorf("status result '%d', status expected '%d'", result, expected)
		}
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
