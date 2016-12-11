package main

import (
	"os"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/the-guitarman/fussball_de_match_list/match_list"
	"log"
	"net/http"
	net_url "net/url"
)

var Version = "0.2"

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 || contains(argsWithoutProg, "-h") || contains(argsWithoutProg, "--help") {
		fmt.Println("-h, --help            show this help")
		fmt.Println("-s, --serve           start server")
		fmt.Println("-u <url>, --url <url> grab and parse a fussball.de team page url and return match list as json")
		fmt.Println("-v, --version         show the version number")
		fmt.Println()
		return
	}

	if contains(argsWithoutProg, "-s") || contains(argsWithoutProg, "--serve") {
		startServer()
		return
	}

	if (contains(argsWithoutProg, "-u") || contains(argsWithoutProg, "--url") && len(argsWithoutProg) == 2) {
		if isUrlValid(argsWithoutProg[1]) {
			//TODO: get match list
		} else {
			fmt.Println("URL is invalid: " + argsWithoutProg[1])
		}
		return
	}

	if contains(argsWithoutProg, "-v") || contains(argsWithoutProg, "--version") {
		fmt.Println("Version: " + Version)
		return
	}
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/match-list", match_list.MatchListHandler)
	http.Handle("/", r)
	fmt.Println("Server: http://localhost:3333/")
	log.Fatal(http.ListenAndServe(":3333", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usage: GET to <your domain:3333>/match-list?url=<encoded url with matchlist of my club at fussball.de>")
}

func isUrlValid(url string) bool {
	result := true
	_, err := net_url.ParseRequestURI(url)
	if err != nil {
	   result = false
	}
	return result
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
