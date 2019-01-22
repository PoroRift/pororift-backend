package dto

type MatchlistDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
	TotalGames int                 `json:"totalGames"`
}
