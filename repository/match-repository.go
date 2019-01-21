package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Pororift/backend/models/dto"
)

func GetMatch(matchId int64) (dto.MatchDTO, error) {
	match := dto.MatchDTO{}
	matchURL := fmt.Sprintf("%s/lol/match/v4/matches/%d", baseURL, matchId)

	req, _ := http.NewRequest("GET", matchURL, nil)
	req.Header.Set("api_key", apiKey)

	res, fetchErr := client.Do(req)
	if fetchErr != nil {
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
