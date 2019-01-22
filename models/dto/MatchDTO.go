package dto

type MatchDTO struct {
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameID                int64                    `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          int64                    `json:"gameDuration"`
	GameCreation          int64                    `json:"gameCreation"`
}
