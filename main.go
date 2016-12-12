package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/the-guitarman/fussball_de_match_list/match_list"
	"log"
	"net/http"
	net_url "net/url"
	"os"
	"strconv"
)

var Version = "0.3"
var DefaultPort = "3333"

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 || contains(argsWithoutProg, "-h") || contains(argsWithoutProg, "--help") {
		fmt.Println("-h, --help                   Shows help information")
		fmt.Println("-p <number>, --port <number> Port number for the server, -s option is required")
		fmt.Println("-s, --serve                  Starts the server mode")
		fmt.Println("-u <url>, --url <url>        Grabs and parses a fussball.de team page url and returns a match list json")
		fmt.Println("-v, --version                Show the version number")
		fmt.Println()
		return
	}

	if contains(argsWithoutProg, "-s") || contains(argsWithoutProg, "--serve") {
		startServer(argsWithoutProg)
		return
	}

	if contains(argsWithoutProg, "-p") || contains(argsWithoutProg, "--port") && len(argsWithoutProg) == 2 {
		fmt.Println("-s option is required")
		return
	}

	if contains(argsWithoutProg, "-u") || contains(argsWithoutProg, "--url") && len(argsWithoutProg) == 2 {
		if isUrlValid(argsWithoutProg[1]) {
			response, err := match_list.GetMatchList(argsWithoutProg[1])
			if err != nil {
				fmt.Println("Error: " + err.Error())
			} else {
				fmt.Println(response);
			}
		} else {
			fmt.Println("URL is invalid: " + argsWithoutProg[1])
		}
		return
	}

	if contains(argsWithoutProg, "-v") || contains(argsWithoutProg, "--version") {
		fmt.Println(os.Args[0] + " v" + Version)
		return
	}
}

func startServer(argsWithoutProg []string) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/match-list", MatchListHandler)
	http.Handle("/", r)

	port := serverPort(argsWithoutProg)
	fmt.Println("Server: http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Usage: GET to <your domain:3333>/match-list?url=<encoded url with matchlist of my club at fussball.de>")
}

func MatchListHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	response, err := match_list.GetMatchList(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		log.Fatal(err)
		return
	}

	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}

func serverPort(argsWithoutProg []string) string {
	result := DefaultPort
	index := portFlagIndex(argsWithoutProg)
	if index > -1 {
		if len(argsWithoutProg) > (index + 1) {
			port := argsWithoutProg[(index + 1)]
			if _, err := strconv.Atoi(port); err == nil {
				result = port
			}
		} else {
			fmt.Println("Port number is missing. Default port " + result + " will be used.")
		}
	}
	return result
}

func portFlagIndex(argsWithoutProg []string) int {
	elementIndex := index(argsWithoutProg, "-p")
	if elementIndex == -1 {
		elementIndex = index(argsWithoutProg, "--port")
	}
	return elementIndex
}

func urlFlagIndex(argsWithoutProg []string) int {
	elementIndex := index(argsWithoutProg, "-u")
	if elementIndex == -1 {
		elementIndex = index(argsWithoutProg, "--url")
	}
	return elementIndex
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

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
