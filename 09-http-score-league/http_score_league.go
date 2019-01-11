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
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		name := request.URL.Path[len("/players/"):]

		switch request.Method {
		case http.MethodPost:
			p.addScore(response, name)
		case http.MethodGet:
			p.showScore(response, name)
		}
	}))

	router.ServeHTTP(response, request)
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
