package main

import (
	"fmt"
	"github.com/jeslopcru/golang-examples/07-http-score-store"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) RecordWin(name string) {
}
func (i *InMemoryPlayerStore) ObtainPlayerScore(name string) int {
	return 123
}

func main() {

	fmt.Println("Server Running in http://localhost:5000")
	server := &http_score_store.PlayerServer{Store: &InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
