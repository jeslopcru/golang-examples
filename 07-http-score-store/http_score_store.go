package http_score_store

import (
	"fmt"
	"net/http"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {

	player := request.URL.Path[len("/players/"):]
	fmt.Fprint(response, ObtainPlayerScore(player))
}

func ObtainPlayerScore(name string) string {
	if name == "Paco" {
		return "20"
	}

	if name == "Manolo" {
		return "35"
	}
	return ""
}
