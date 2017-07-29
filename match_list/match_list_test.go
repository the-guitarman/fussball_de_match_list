package match_list

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"encoding/json"
)

func TestGetMatchList(t *testing.T) {
	response, err := GetMatchList("file://../test.html")
	if err != nil {
		panic(err)
	}

	var match_list MatchList
	bytes := []byte(response)
	err = json.Unmarshal(bytes, &match_list)
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
