package http_score_store

import (
	"fmt"
	"net/http"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	ObtainPlayerScore(name string) int
}

// PlayerServer is a HTTP interface for player information
type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	player := request.URL.Path[len("/players/"):]

	score := p.Store.ObtainPlayerScore(player)
	if 0 == score {
		response.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(response, score)
}
