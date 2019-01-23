package services

import (
	"github.com/Pororift/backend/models/dto"
	"github.com/Pororift/backend/repository"
)

// GetMatch a service that will fetch a given match
// from the riot api and transform it to an appropriate data object
// that can be used in our side.
func GetMatch(matchID int64) (*dto.MatchDTO, error) {
	match, err := repository.GetMatch(matchID)
	return match, err
}

// GetMatchHistory is a service that will fetch a list of matches
// a given accountID has had in the past
func GetMatchHistory(accountID string) (*dto.MatchlistDTO, error) {
	matchList, err := repository.GetMatchHistory(accountID)
	return matchList, err
}
