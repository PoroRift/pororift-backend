package services

import (
	"github.com/Pororift/backend/models/dto"
	"github.com/Pororift/backend/repository"
)

func GetMatch(matchId int64) (dto.MatchDTO, error) {
	match, err := repository.GetMatch(matchId)
	return match, err
}
