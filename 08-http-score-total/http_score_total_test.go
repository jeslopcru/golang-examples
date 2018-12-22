package http_score_total

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) ObtainPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestHttpScore(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Paco":   20,
			"Manolo": 35,
		}, nil,
	}
	server := &PlayerServer{&store}

	t.Run("return Paco's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Paco")
		server.ServeHTTP(response, request)

		assertStatus(response.Code, http.StatusOK, t)
		assertResponseBody(t, "20", response.Body.String())

	})

	t.Run("return Manolo's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Manolo")
		server.ServeHTTP(response, request)

		assertStatus(response.Code, http.StatusOK, t)
		assertResponseBody(t, "35", response.Body.String())
	})

	t.Run("return 404 when not found a player", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := createNewRequest("Juan")
		server.ServeHTTP(response, request)

		assertStatus(response.Code, http.StatusNotFound, t)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{}, nil,
	}
	server := &PlayerServer{&store}

	t.Run("it record a win when POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/Luis", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(response.Code, http.StatusAccepted, t)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})
}

func assertStatus(result int, expected int, t *testing.T) {
	if result != expected {
		t.Errorf("wrong status status result '%d', status expected '%d'", result, expected)
	}
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
