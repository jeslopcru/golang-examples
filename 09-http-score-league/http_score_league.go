package http_score_league

import (
	"fmt"
	"net/http"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	ObtainPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a HTTP interface for player information

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

func PlayerServerMaster(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}

	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	p.router.ServeHTTP(response, request)
}

func (p *PlayerServer) leagueHandler(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
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
