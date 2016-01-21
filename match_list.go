package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Match struct {
	Start_at string `json:"start_at"`
	Team_one string `json:"home"`
	Team_two string `json:"guest"`
}

type MatchList struct {
	Team_name string  `json:"team_name"`
	Matches   []Match `json:"matches"`
}

func MatchListHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	doc, err := goquery.NewDocument(url)
	//fmt.Println(doc)
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	var headline string
	var match_list MatchList = MatchList{Team_name: doc.Find("h2").First().Text()}
	var next_match_time_row_index = 0
	var next_team_row_index = 2

	lines := doc.Find("#team-matchplan-table tbody tr")
	lines.Each(func(i int, s *goquery.Selection) {
		if i == next_match_time_row_index { // headline
			headline = strings.Trim(strings.Split(s.Find("td").Text(), "|")[0], " ")
			next_match_time_row_index += 3
			return
		}

		if i == next_team_row_index { // team names
			m := Match{
				Start_at: headline,
				Team_one: s.Find("td.column-club .club-name").First().Text(),
				Team_two: s.Find("td.column-club .club-name").Last().Text(),
			}
			match_list.Matches = append(match_list.Matches, m)
			next_team_row_index += 3
			return
		}
	})

	resp, _ := json.MarshalIndent(match_list, "", "  ")

	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(resp))
}
