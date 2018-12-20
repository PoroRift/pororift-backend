package data

import "fmt"

type (
	Matches struct {
		Matches    []Match `json: matches`
		StartIndex int     `json: startindex`
		EndIndex   int     `json: endIndex`
		TotalGames int     `json: totalGames`
	}

	Match struct {
		PlatformId string `json: platformId`
		GameId     int    `json: gameId`
		Champion   int    `json: champion`
		Queue      int    `json: queue`
		Season     int    `json:season`
		Timestamp  int    `json:timestamp`
		Role       string `json: role`
		Lane       string `json:lane`
	}
)

func (s *Matches) String() string {
	return fmt.Sprintf("%+v\n", *s)
}
