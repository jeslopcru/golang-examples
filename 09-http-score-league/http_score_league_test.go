package http_score_league

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
	server := PlayerServerMaster(&store)

	t.Run("return Paco's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest("Paco")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, "20", response.Body.String())

	})

	t.Run("return Manolo's score", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest("Manolo")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, "35", response.Body.String())
	})

	t.Run("return 404 when not found a player", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := newGetScoreRequest("Juan")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{}, nil,
	}
	server := PlayerServerMaster(&store)

	t.Run("it record a win when POST", func(t *testing.T) {
		player := "Luis"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not Store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func assertStatus(t *testing.T, result int, expected int) {
	if result != expected {
		t.Errorf("wrong status status result '%d', status expected '%d'", result, expected)
	}
}

func assertResponseBody(t *testing.T, expected string, result string) {
	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := PlayerServerMaster(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
