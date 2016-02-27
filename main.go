package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/match-list", MatchListHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3333", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usage: GET to <your domain:3333>/match-list?url=<encoded url with matchlist of my club at fussball.de>")
}
