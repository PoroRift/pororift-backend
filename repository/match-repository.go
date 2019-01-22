package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Pororift/backend/models/dto"
	"github.com/labstack/echo"
)

func GetMatch(matchId int64) (dto.MatchDTO, error) {
	var match dto.MatchDTO
	e := echo.New()

	matchURL := getBaseUrl(fmt.Sprintf("/lol/match/v4/matches/%d", matchId))

	req, _ := http.NewRequest("GET", matchURL, nil)

	res, fetchErr := client.Do(req)
	if fetchErr != nil || res.StatusCode != 200 {
		return match, fetchErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return match, readErr
	}

	jsonErr := json.Unmarshal(body, &match)

	if jsonErr != nil {
		return match, jsonErr
	}
	return match, nil
}

func GetMatchHistory(accountID string) (dto.MatchlistDTO, error) {
	matchListURL := getBaseUrl(fmt.Sprintf("/lol/match/v4/matchlists/by-account/%s", accountId))
	
	req, _ := http.NewRequest("GET", matchListURL, nil)

	res, fetch
}


