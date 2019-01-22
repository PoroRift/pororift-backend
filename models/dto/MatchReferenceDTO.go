package dto

type MatchReferenceDTO struct {
	PlatformID string `json:"platformId"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Season     int    `json:"season"`
	Timestamp  int64  `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}
