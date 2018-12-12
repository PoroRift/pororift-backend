package lol

import "fmt"

// type Player struct {
// 	Stat
// }

// type Stat struct {
// 	id string
// }

type (
	Stat struct {
		Id           string `json:"id"`
		AccountId    string `json:"accountId"`
		Puuid        string `json:"puuid"`
		Name         string `json:"name"`
		ProfileIcon  int    `json:"profileIconId"`
		RevisionDate int    `json:"revisionDate"`
		Lvl          int    `json:"summonerLevel"`
	}

	Player struct {
		Stat
	}
)

// "ToString" for type Stat
func (s *Stat) String() string {
	// Print the name of the field along with value
	return fmt.Sprintf("%+v\n", *s)
}
