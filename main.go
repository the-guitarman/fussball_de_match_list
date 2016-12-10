package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/the-guitarman/fussball_de_match_list/match_list"
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/match-list", match_list.MatchListHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3333", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usage: GET to <your domain:3333>/match-list?url=<encoded url with matchlist of my club at fussball.de>")
}
