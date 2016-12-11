package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	//"fmt"
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

	assert.Equal(t, 9, len(match_list.Matches))
	assert.Equal(t, "TSV Einheit Claußnitz", match_list.Team_name)

	match := match_list.Matches[0]
	assert.Equal(t, "2016-11-20T14:00:00+01:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "TSV Dittersbach", match.Team_two)

	match = match_list.Matches[1]
	assert.Equal(t, "2016-11-27T14:00:00+01:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "BSC Motor Rochlitz 2", match.Team_two)

	match = match_list.Matches[2]
	assert.Equal(t, "2016-12-04T14:00:00+01:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "FSV Zettlitz", match.Team_one)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_two)

	match = match_list.Matches[3]
	assert.Equal(t, "2017-04-02T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "SpG Königshain-Wiederau/​Wechselburg", match.Team_two)

	match = match_list.Matches[4]
	assert.Equal(t, "2017-04-09T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "SC 1999 Altmittweida", match.Team_one)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_two)

	match = match_list.Matches[5]
	assert.Equal(t, "2017-04-23T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "SV 05 Hartmannsdorf 2", match.Team_two)

	match = match_list.Matches[6]
	assert.Equal(t, "2017-04-30T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "LSV Sachsenburg", match.Team_one)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_two)

	match = match_list.Matches[7]
	assert.Equal(t, "2017-05-07T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_one)
	assert.Equal(t, "SV 94 Geringswalde/​Schweikershain", match.Team_two)

	match = match_list.Matches[8]
	assert.Equal(t, "2017-05-14T15:00:00+02:00", match.Start_at)
	assert.Equal(t, "2.Kreisliga (B)", match.Competition)
	assert.Equal(t, "TV Vater Jahn Burgstädt", match.Team_one)
	assert.Equal(t, "TSV Einheit Claußnitz", match.Team_two)
}
