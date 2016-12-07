package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

type Match struct {
	Start_at    string `json:"start_at"`
	Competition string `json:"competition"`
	Team_one    string `json:"home"`
	Team_two    string `json:"guest"`
}

type MatchList struct {
	Team_name string  `json:"team_name"`
	Matches   []Match `json:"matches"`
}

func MatchListHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	doc, err := goquery.NewDocument(url)
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	var headline string
	var start_at string
	var competition string
	var match_list MatchList = MatchList{Team_name: findTeamName(doc)}
	var next_match_time_row_index = 0
	var next_team_row_index = 2

	lines := doc.Find("#team-matchplan-table tbody tr")
	lines.Each(func(i int, s *goquery.Selection) {
		if i == next_match_time_row_index { // headline
			headline = s.Find("td").Text()
			start_at = strings.Trim(strings.Split(headline, "|")[0], " ")
			competition = strings.Trim(strings.Split(headline, "|")[1], " ")
			next_match_time_row_index += 3
			return
		}

		if i == next_team_row_index { // team names
			m := Match{
				Start_at:    parseDate(start_at),
				Competition: competition,
				Team_one:    s.Find("td.column-club .club-name").First().Text(),
				Team_two:    s.Find("td.column-club .club-name").Last().Text(),
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

func findTeamName(d *goquery.Document) (n string) {
	n = d.Find("h2").First().Text()
	return
}

func parseDate(date string) string {
	loc, _ := time.LoadLocation("Europe/Berlin")
	//errorCheck(err)
	layout := "Sonntag, 02.01.2006 - 15:04 Uhr"
	myTime, _ := time.ParseInLocation(layout, date, loc)
	//errorCheck(err)
	return myTime.Format("2006-01-02T15:04:05") + timeZone(myTime, loc)
}

func timeZone(myTime time.Time, loc *time.Location) string {
	timeZone := "+01:00"
	_, timeOffset := myTime.Zone()
	zw, winterOffset := time.Date(myTime.Year(), 1, 1, 0, 0, 0, 0, loc).Zone()
	zs, summerOffset := time.Date(myTime.Year(), 6, 1, 0, 0, 0, 0, loc).Zone()

	if winterOffset > summerOffset {
		winterOffset, summerOffset = summerOffset, winterOffset
		zw, zs = zs, zw
	}

	if winterOffset != summerOffset { // the location has daylight saving
		if timeOffset != winterOffset {
			timeZone = "+02:00"
		}
	}
	return timeZone
}
