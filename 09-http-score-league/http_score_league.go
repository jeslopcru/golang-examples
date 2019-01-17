package http_score_league

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Name  string
	Score int
}

// PlayerStore stores score information about players
type PlayerStore interface {
	ObtainPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func PlayerServerMaster(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(p.obtainLeagueList())
	response.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) obtainLeagueList() []Player {
	return []Player{
		{"Paco", 20},
	}
}

func (p *PlayerServer) playersHandler(response http.ResponseWriter, request *http.Request) {
	name := request.URL.Path[len("/players/"):]

	switch request.Method {
	case http.MethodPost:
		p.addScore(response, name)
	case http.MethodGet:
		p.showScore(response, name)
	}
}

func (p *PlayerServer) addScore(response http.ResponseWriter, name string) {
	p.store.RecordWin(name)
	response.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showScore(response http.ResponseWriter, name string) {

	score := p.store.ObtainPlayerScore(name)
	if 0 == score {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}
