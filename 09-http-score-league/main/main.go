package main

import (
	"fmt"
	"github.com/jeslopcru/golang-examples/09-http-score-league"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server Running in http://localhost:5000")
	server := &http_score_league.PlayerServer{Store: &http_score_league.InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
