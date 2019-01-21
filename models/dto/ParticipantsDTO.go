package dto

type ParticipantDTO struct {
	Stats                     ParticipantStatsDTO    `json:"stats"`
	ParticipantID             int                    `json:"participantId"`
	Runes                     []RuneDTO              `json:"runes"`
	Timeline                  ParticipantTimelineDTO `json:"timeline"`
	TeamID                    int                    `json:"teamId"`
	Spell2ID                  int                    `json:"spell2Id"`
	Masteries                 []MasteryDTO           `json:"masteries"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
	Spell1ID                  int                    `json:"spell1Id"`
	ChampionID                int                    `json:"championID"`
}
