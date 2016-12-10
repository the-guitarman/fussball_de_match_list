package match_list

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
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

func init() {

}

func MatchListHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	doc := goqueryDocument(url, w);

	matchList := parseMatchList(doc)
	response := matchListToJson(matchList, w)

	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}

func goqueryDocument(url string, w http.ResponseWriter) (doc *goquery.Document) {
	if localUrl(url) {
		filePath := strings.Replace(url, "file://", "", 1)
		doc, _ = goqueryDocumentFromFile(filePath, w)
	} else {
		doc, _ = goqueryDocumentFromUrl(url, w)
	}
	return
}

func goqueryDocumentFromFile(filePath string, w http.ResponseWriter) (doc *goquery.Document, err error) {
	fileContent := readFile(filePath, w)
	stringReader := strings.NewReader(fileContent)
	doc, err = goquery.NewDocumentFromReader(stringReader)
	handleError(err, w)
	return
}

func readFile(filePath string, w http.ResponseWriter) (fileContent string) {
	byteContent, err := ioutil.ReadFile(filePath)
	handleError(err, w)
	fileContent = string(byteContent)
	return
}

func goqueryDocumentFromUrl(url string, w http.ResponseWriter) (doc *goquery.Document, err error) {
	doc, err = goquery.NewDocument(url)
	handleError(err, w)
	return
}

func localUrl(url string) bool {
	ret := false
	if strings.HasPrefix(url, "file://") {
		ret = true
	}
	return ret
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		//log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
}

func parseMatchList(doc *goquery.Document) (matchList MatchList) {
	matchList = MatchList{Team_name: findTeamName(doc)}

	var headline string
	var start_at string
	var competition string
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
			matchList.Matches = append(matchList.Matches, m)
			next_team_row_index += 3
			return
		}
	})
	return
}

func matchListToJson(matchList MatchList, w http.ResponseWriter) (string) {
	response, err := json.MarshalIndent(matchList, "", "  ")
	handleError(err, w)
	return string(response)
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
