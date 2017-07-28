package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"encoding/json"
	"github.com/the-guitarman/fussball_de_match_list/match_list"
	"io/ioutil"
	"net/http"
	"time"
)

func TestAllRoutes(t *testing.T) {
	go startServer([]string{"-s"})

	client := &http.Client{Timeout: 1 * time.Second}

	homeRouteTest(t, client)
	matchListTest(t, client)
}

func homeRouteTest(t *testing.T, client *http.Client) {
	request, _ := http.NewRequest("GET", "http://localhost:3333/", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Usage: GET to <your domain:3333>/match-list?url=<encoded url with matchlist of my club at fussball.de>", string(body))
}

func matchListTest(t *testing.T, client *http.Client) {
	request, _ := http.NewRequest("GET", "http://localhost:3333/match-list?url=file://test.html", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	//assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var match_list match_list.MatchList
	err = json.Unmarshal(body, &match_list)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 10, len(match_list.Matches))
	assert.Equal(t, "TSV Einheit Claußnitz", match_list.Team_name)

	match := match_list.Matches[0]
	assert.Equal(t, "2017-07-30T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "Kreisfreundschaftsspiele", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "FC Wacker 90 Wittgensdorf", match.Team_two)
}
