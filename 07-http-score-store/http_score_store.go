package http_score_store

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
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	name := request.URL.Path[len("/players/"):]
	switch request.Method {
	case http.MethodPost:
		p.addScore(response, name)
	case http.MethodGet:
		p.showScore(response, name)
	}
}

func (p *PlayerServer) addScore(response http.ResponseWriter, name string) {
	p.Store.RecordWin(name)
	response.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showScore(response http.ResponseWriter, name string) {

	score := p.Store.ObtainPlayerScore(name)
	if 0 == score {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}
