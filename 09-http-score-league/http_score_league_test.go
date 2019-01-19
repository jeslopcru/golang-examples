package http_score_league

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) ObtainPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestHttpScore(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Paco":   20,
			"Manolo": 35,
		}, nil, nil,
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
		map[string]int{}, nil, nil,
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
	t.Run("it returns league list as JSON", func(t *testing.T) {
		expectedLeague := []Player{
			{"Juan", 7},
			{"Manolo", 5},
			{"Paco", 2},
		}
		store := StubPlayerStore{nil, nil, expectedLeague}
		server := PlayerServerMaster(&store)

		request := leagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		result := obtainLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(response, t)
		assertLeague(result, expectedLeague, t)
	})
}

func assertContentType(response *httptest.ResponseRecorder, t *testing.T) {
	if response.Header().Get("content-type") != "application/json" {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}

func assertLeague(result []Player, expectedLeague []Player, t *testing.T) {
	if !reflect.DeepEqual(result, expectedLeague) {
		t.Errorf("got %v want %v", result, expectedLeague)
	}
}

func obtainLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}

func leagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}
