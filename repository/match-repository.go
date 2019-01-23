package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Pororift/backend/models/dto"
)

// GetMatch gets the match given a match id
func GetMatch(matchID int64) (*dto.MatchDTO, error) {

	matchURL := getBaseURL(fmt.Sprintf("/lol/match/v4/matches/%d", matchID))

	req, reqErr := http.NewRequest("GET", matchURL, nil)

	obj, callErr := callHandler(req, reqErr)

	if obj == nil || callErr != nil {
		return nil, callErr
	}
	match, ok := (*obj).(dto.MatchDTO)
	if ok {
		return &match, nil
	}
	return nil, errors.New("did not receive a match object")
}

// GetMatchHistory gets a list of matches given an account id
func GetMatchHistory(accountID string) (*dto.MatchlistDTO, error) {

	matchListURL := getBaseURL(fmt.Sprintf("/lol/match/v4/matchlists/by-account/%s", accountID))

	req, reqErr := http.NewRequest("GET", matchListURL, nil)

	obj, callErr := callHandler(req, reqErr)

	if obj == nil || callErr != nil {
		return nil, callErr
	}
	matchList, ok := (*obj).(dto.MatchlistDTO)
	if ok {
		return &matchList, nil
	}
	return nil, errors.New("did not receive list of matches")
}
