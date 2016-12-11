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

var Version = "0.2"
var DefaultPort = "3333"

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 || contains(argsWithoutProg, "-h") || contains(argsWithoutProg, "--help") {
		fmt.Println("-h, --help                   show this help")
		fmt.Println("-p <number>, --port <number> port number for the server, -s option is required")
		fmt.Println("-s, --serve                  start server")
		fmt.Println("-u <url>, --url <url>        grab and parse a fussball.de team page url and return match list as json")
		fmt.Println("-v, --version                show the version number")
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

func startServer(argsWithoutProg []string) {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/match-list", match_list.MatchListHandler)
	http.Handle("/", r)

	port := serverPort(argsWithoutProg)
	fmt.Println("Server: http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func serverPort(argsWithoutProg []string) string {
	result := DefaultPort
	index := portFlagIndex(argsWithoutProg)
	if index > -1 {
		if len(argsWithoutProg) > (index+1) {
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

func index(s []string, e string) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
